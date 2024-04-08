package main

import (
	"log/slog"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/elastic/go-elasticsearch/v8"

	"github.com/tuannkhoi/sport-data-feed/config"
	"github.com/tuannkhoi/sport-data-feed/service"
)

func main() {
	logger := slog.Default()

	cfg := config.NewConfig()

	// extra config for consumer
	(*cfg.KafkaConfigMap)["group.id"] = "go-group-1"
	(*cfg.KafkaConfigMap)["auto.offset.reset"] = "earliest"

	dynamoDBClient := dynamodb.NewFromConfig(*cfg.AWSConfig)

	esClient, err := elasticsearch.NewTypedClient(*cfg.ElasticsearchConfig)

	sdc, err := service.NewSportDataConsumer(cfg, logger, dynamoDBClient, esClient)
	if err != nil {
		logger.Error("Failed to create SportDataConsumer: ", err)

		return
	}

	sdc.Consume()
}
