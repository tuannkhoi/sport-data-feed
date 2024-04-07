package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"

	"github.com/tuannkhoi/sport-data-feed/sports"
	"github.com/tuannkhoi/sport-data-feed/utils"
)

func main() {
	// sets the consumer group ID and offset
	conf := utils.ReadConfig()
	conf["group.id"] = "go-group-1"
	conf["auto.offset.reset"] = "earliest"
	topic := "football-match-new"

	// creates a new consumer and subscribes to your topic
	consumer, err := kafka.NewConsumer(&conf)
	if err != nil {
		log.Fatalf("Failed to create consumer: %s\n", err)
	}

	if err := consumer.SubscribeTopics([]string{topic}, nil); err != nil {
		log.Fatalf("Failed to subscribe to topic: %s\n", err)
	}

	run := true

	for run {
		// consumes messages from the subscribed topic and prints them to the console
		e := consumer.Poll(1000)
		switch ev := e.(type) {
		case *kafka.Message:
			// application-specific processing
			fmt.Printf("Consumed event from topic %s: key = %-10s value = %s\n",
				*ev.TopicPartition.Topic, string(ev.Key), "see below")

			newFootBallMatch := new(sports.Match)

			if err := json.Unmarshal(ev.Value, newFootBallMatch); err != nil {
				log.Fatalln(err)
			}

			fmt.Printf("Match ID: %s\n", newFootBallMatch.ID)
			fmt.Printf("Home Team: %s\n", newFootBallMatch.HomeTeam)
			fmt.Printf("Away Team: %s\n", newFootBallMatch.AwayTeam)
			fmt.Printf("Stadium: %s\n", newFootBallMatch.Stadium)
			fmt.Printf("Round: %s\n", newFootBallMatch.Round)
			fmt.Printf("Competition: %s\n", newFootBallMatch.Competition)
			fmt.Printf("Kick Off: %s\n", newFootBallMatch.KickOff)
			fmt.Printf("Note: %s\n", newFootBallMatch.Note)
			fmt.Println()
		case kafka.Error:
			_, _ = fmt.Fprintf(os.Stderr, "%% Error: %v\n", ev)

			run = false
		}
	}

	if err := consumer.Close(); err != nil {
		log.Fatalf("Failed to close consumer: %s\n", err)
	}
}
