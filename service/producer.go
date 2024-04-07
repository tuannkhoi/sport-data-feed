package service

import (
	"encoding/json"
	"errors"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"

	"github.com/tuannkhoi/sport-data-feed/sports"
	"github.com/tuannkhoi/sport-data-feed/utils"
)

type SportDataProducer struct {
	Producer *kafka.Producer
	Log      *slog.Logger
}

// NewSportDataProducer creates a new SportDataProducer instance.
func NewSportDataProducer(logger *slog.Logger) (*SportDataProducer, error) {
	conf := utils.ReadConfig()

	producer, err := kafka.NewProducer(&conf)
	if err != nil {
		return nil, errors.New("Failed to create Producer: " + err.Error())
	}

	return &SportDataProducer{
		Producer: producer,
		Log:      logger,
	}, nil
}

// ProduceNewFootballMatch produces a new football match every 3 seconds.
func (sdp *SportDataProducer) ProduceNewFootballMatch() {
	topic := "football-match-new"

	ticker := time.NewTicker(3 * time.Second)
	defer ticker.Stop()

	sigCh := make(chan os.Signal, 1)
	defer close(sigCh)

	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

produceLoop:
	for {
		select {
		case <-sigCh:
			sdp.Log.Info("Received signal to close the producer. Closing...")

			// send any outstanding or buffered messages to the Kafka broker and close the connection
			sdp.Producer.Flush(15 * 1000)
			sdp.Producer.Close()

			break produceLoop
		case <-ticker.C:
			footballMatch := sports.NewFootballMatch()

			bytes, err := json.Marshal(footballMatch)
			if err != nil {
				sdp.Log.Warn("Failed to marshal football match: " + err.Error())
			}

			// produces a sample message to the user-created topic
			if err := sdp.Producer.Produce(&kafka.Message{
				TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
				Key:            []byte(footballMatch.ID.String()),
				Value:          bytes,
			}, nil); err != nil {
				sdp.Log.Warn("Failed to produce message: " + err.Error())
			}
		}
	}
}

// Monitor handle message delivery reports and possibly other event types (errors, stats, etc.,).
func (sdp *SportDataProducer) Monitor() {
	for e := range sdp.Producer.Events() {
		switch ev := e.(type) {
		case *kafka.Message:
			if ev.TopicPartition.Error != nil {
				sdp.Log.Warn("Failed to deliver message: " + ev.TopicPartition.Error.Error())
			} else {
				sdp.Log.Info("Produced event to topic " + *ev.TopicPartition.Topic)
			}
		}
	}
}
