package main

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"log"
)

const TopicName = "ExampleTopic"

func main() {
	forever := make(chan bool)

	producer := NewKafkaProducer()
	_ = Publish("HelloWorld Go", TopicName, producer)

	<-forever
}

func NewKafkaProducer() *kafka.Producer {
	configMap := kafka.ConfigMap{
		"bootstrap.servers": "host.docker.internal:9092",
	}

	producer, err := kafka.NewProducer(&configMap)
	if err != nil {
		log.Fatal(err.Error())
	}

	return producer
}

func Publish(msg string, topic string, producer *kafka.Producer) error {
	message := kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &topic,
			Partition: kafka.PartitionAny,
		},
		Value: []byte(msg),
	}

	err := producer.Produce(&message, nil)
	if err != nil {
		log.Println("error", err.Error())
	}

	return nil
}
