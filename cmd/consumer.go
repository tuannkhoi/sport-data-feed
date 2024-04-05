package main

import (
	"fmt"
	"os"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"

	"github.com/tuannkhoi/sport-data-feed/utils"
)

func main() {
	// sets the consumer group ID and offset
	conf := utils.ReadConfig()
	conf["group.id"] = "go-group-1"
	conf["auto.offset.reset"] = "earliest"
	topic := "topic_0"

	// creates a new consumer and subscribes to your topic
	consumer, _ := kafka.NewConsumer(&conf)
	consumer.SubscribeTopics([]string{topic}, nil)

	run := true

	var numMessagesReceived int

	for run {
		// consumes messages from the subscribed topic and prints them to the console
		e := consumer.Poll(1000)
		switch ev := e.(type) {
		case *kafka.Message:
			// application-specific processing
			fmt.Printf("Consumed event from topic %s: key = %-10s value = %s\n",
				*ev.TopicPartition.Topic, string(ev.Key), string(ev.Value))
			numMessagesReceived++
			fmt.Println("Message number", numMessagesReceived)
		case kafka.Error:
			fmt.Fprintf(os.Stderr, "%% Error: %v\n", ev)
			run = false
		}
	}

	// closes the consumer connection
	consumer.Close()

}
