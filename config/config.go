package config

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/spf13/viper"
)

type Config struct {
	KafkaConfigMap *kafka.ConfigMap
	AWSConfig      *aws.Config
}

func init() {
	viper.AddConfigPath("deploy")
	viper.SetConfigName("config")
	viper.SetConfigType("toml")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalln(fmt.Sprintf("Failed to read config file: %s", err))
	}
}

func NewConfig() *Config {
	return &Config{
		KafkaConfigMap: readKafkaConfig(),
		AWSConfig:      readAWSConfig(),
	}
}

func readKafkaConfig() *kafka.ConfigMap {
	kafkaConfigMap := make(kafka.ConfigMap)

	stringKeys := []string{
		"bootstrap.servers",
		"security.protocol",
		"sasl.mechanism",
		"sasl.username",
		"sasl.password",
	}

	for _, key := range stringKeys {
		kafkaConfigMap[key] = viper.GetString(fmt.Sprintf("kafka.%s", key))
	}

	intKeys := []string{"session.timeout.ms"}

	for _, key := range intKeys {
		kafkaConfigMap[key] = viper.GetInt(fmt.Sprintf("kafka.%s", key))
	}

	return &kafkaConfigMap
}

func readAWSConfig() *aws.Config {
	return &aws.Config{
		Region: viper.GetString("aws.region"),
		Credentials: aws.CredentialsProviderFunc(func(_ context.Context) (aws.Credentials, error) {
			return aws.Credentials{
				AccessKeyID:     viper.GetString("aws.access_key_id"),
				SecretAccessKey: viper.GetString("aws.secret_access_key"),
			}, nil
		}),
	}
}
