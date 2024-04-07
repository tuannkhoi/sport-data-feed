package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"

	"github.com/tuannkhoi/sport-data-feed/sports"
	"github.com/tuannkhoi/sport-data-feed/utils"
)

type SportDataConsumer struct {
	Consumer *kafka.Consumer
	Log      *slog.Logger
}

// NewSportDataConsumer creates a new SportDataConsumer instance.
func NewSportDataConsumer(logger *slog.Logger) (*SportDataConsumer, error) {
	conf := utils.ReadConfig()
	conf["group.id"] = "go-group-1"
	conf["auto.offset.reset"] = "earliest"

	consumer, err := kafka.NewConsumer(&conf)
	if err != nil {
		return nil, errors.New("Failed to create Consumer: " + err.Error())
	}

	return &SportDataConsumer{
		Consumer: consumer,
		Log:      logger,
	}, nil
}

func (sdc *SportDataConsumer) Consume() {
	if err := sdc.Consumer.SubscribeTopics([]string{"football-match-new"}, nil); err != nil {
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
			// consumes messages from the subscribed topic and prints them to the console
			e := sdc.Consumer.Poll(1000)
			switch ev := e.(type) {
			case *kafka.Message:
				fmt.Printf("Consumed event from topic %s: key = %-10s value = %s\n\n",
					*ev.TopicPartition.Topic, string(ev.Key), "see below")

				newFootBallMatch := new(sports.FootballMatch)

				if err := json.Unmarshal(ev.Value, newFootBallMatch); err != nil {
					log.Fatalln(err)
				}

				fmt.Printf("FootballMatch ID: %s\n", newFootBallMatch.ID)
				fmt.Printf("Home Team: %s\n", newFootBallMatch.HomeTeam.Name)
				fmt.Printf("Away Team: %s\n", newFootBallMatch.AwayTeam.Name)
				fmt.Printf("Stadium: %s\n", newFootBallMatch.Stadium)
				fmt.Printf("Round: %d\n", newFootBallMatch.Round)
				fmt.Printf("Competition: %s\n", newFootBallMatch.Competition)
				fmt.Printf("Country: %s\n", newFootBallMatch.Country)
				fmt.Printf("Kick Off: %s\n", newFootBallMatch.KickOff)
				fmt.Println()

			case kafka.Error:
				sdc.Log.Error("Error polling for message: " + ev.Error())

				break consumeLoop
			}
		}
	}
}
