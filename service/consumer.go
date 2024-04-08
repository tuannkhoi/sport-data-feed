package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"

	"github.com/tuannkhoi/sport-data-feed/config"
	"github.com/tuannkhoi/sport-data-feed/sports"
)

type SportDataConsumer struct {
	Consumer       *kafka.Consumer
	Log            *slog.Logger
	DynamoDBClient *dynamodb.Client
}

// NewSportDataConsumer creates a new SportDataConsumer instance.
func NewSportDataConsumer(
	cfg *config.Config,
	logger *slog.Logger,
	dynamoDBClient *dynamodb.Client,
) (*SportDataConsumer, error) {
	consumer, err := kafka.NewConsumer(cfg.KafkaConfigMap)
	if err != nil {
		return nil, errors.New("Failed to create Consumer: " + err.Error())
	}

	return &SportDataConsumer{
		Consumer:       consumer,
		Log:            logger,
		DynamoDBClient: dynamoDBClient,
	}, nil
}

func (sdc *SportDataConsumer) Consume() {
	if err := sdc.Consumer.SubscribeTopics([]string{sports.TopicNewFootballMatch}, nil); err != nil {
		log.Fatalf("Failed to subscribe to topic: %s\n", err)
	}

	sigCh := make(chan os.Signal, 1)
	defer close(sigCh)

	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

consumeLoop:
	for {
		select {
		case <-sigCh:
			sdc.Log.Info("Received signal to close the consumer. Closing...")

			if err := sdc.Consumer.Close(); err != nil {
				sdc.Log.Warn("Failed to close consumer: " + err.Error())
			}

			break consumeLoop
		default:
			msg, err := sdc.Consumer.ReadMessage(1000 * time.Millisecond)
			if err != nil {
				if err.Error() == kafka.ErrTimedOut.String() {
					continue
				}

				sdc.Log.Error("Error reading message: " + err.Error())
			}

			fmt.Printf("Consumed event from topic %s: key = %-10s value = %s\n\n",
				*msg.TopicPartition.Topic, string(msg.Key), "see below")

			switch *msg.TopicPartition.Topic {
			case sports.TopicNewFootballMatch:
				fm := new(sports.FootballMatch)

				if err := json.Unmarshal(msg.Value, fm); err != nil {
					sdc.Log.Error("Failed to unmarshal football match: " + err.Error())

					continue
				}

				fmt.Println(msg)

				if err := sdc.HandleNewFootballMatch(fm); err != nil {
					sdc.Log.Error("Failed to handle new football match: " + err.Error())
				}
			}
		}
	}
}

func (sdc *SportDataConsumer) HandleNewFootballMatch(fm *sports.FootballMatch) error {
	fmt.Printf("FootballMatch ID: %s\n", fm.ID)
	fmt.Printf("Home Team: %s\n", fm.HomeTeam.Name)
	fmt.Printf("Away Team: %s\n", fm.AwayTeam.Name)
	fmt.Printf("Stadium: %s\n", fm.Stadium)
	fmt.Printf("Round: %d\n", fm.Round)
	fmt.Printf("Competition: %s\n", fm.Competition)
	fmt.Printf("Country: %s\n", fm.Country)
	fmt.Printf("Kick Off: %s\n", fm.KickOff.String())

	fmt.Println()

	// add it to a batch, regularly flush the batch to DynamoDB
	if _, err := sdc.DynamoDBClient.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName: aws.String("FootballMatches"),
		Item:      fm.ToDynamoDBItem(),
	}); err != nil {
		return errors.New("Failed to put item to DynamoDB: " + err.Error())
	}

	sdc.Log.Info(fmt.Sprintf("Successfully added new football match to DynamoDB: %s\n", fm.ID.String()))

	return nil
}
