package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"

	"github.com/tuannkhoi/sport-data-feed/sports"
	"github.com/tuannkhoi/sport-data-feed/utils"
)

func main() {
	// creates a new producer instance
	conf := utils.ReadConfig()
	p, _ := kafka.NewProducer(&conf)
	topic := "football-match-new"

	// go-routine to handle message delivery reports and
	// possibly other event types (errors, stats, etc)
	go func() {
		for e := range p.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Printf("Failed to deliver message: %v\n", ev.TopicPartition)
				} else {
					fmt.Printf("Produced event to topic %s: key = %-10s value = %s\n",
						*ev.TopicPartition.Topic, string(ev.Key), string(ev.Value))
				}
			}
		}
	}()

	ticker := time.NewTicker(3 * time.Second)
	defer ticker.Stop()

	done := make(chan struct{})

	go func() {
		for {
			select {
			case <-done:
				// send any outstanding or buffered messages to the Kafka broker and close the connection
				p.Flush(15 * 1000)
				p.Close()
			case <-ticker.C:
				footballMatch := sports.NewMatch()

				bytes, err := json.Marshal(footballMatch)
				if err != nil {
					log.Fatalln(err)
				}

				// produces a sample message to the user-created topic
				p.Produce(&kafka.Message{
					TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
					Key:            []byte(footballMatch.ID.String()),
					Value:          bytes,
				}, nil)
			}
		}
	}()

	fmt.Println("press Enter to stop")
	_, _ = fmt.Scanln()
	done <- struct{}{}
}
