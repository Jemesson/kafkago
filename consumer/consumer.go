package main

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"log"
)

const TopicName = "ExampleTopic"

func main() {

	ch := make(chan kafka.Message)

	go consume(ch)

	for msg := range ch {
		fmt.Println("Getting messsage", string(msg.Value))
	}
}

func consume(ch chan kafka.Message) {

	configMap := kafka.ConfigMap{
		"bootstrap.servers": "host.docker.internal:9092",
		"group.id":          "mygroup",
	}

	consumer, err := kafka.NewConsumer(&configMap)
	if err != nil {
		log.Fatal(err.Error())
	}

	consumer.Subscribe(TopicName, nil)
	for {
		msg, err := consumer.ReadMessage(-1)
		if err != nil {
			ch <- *msg
		}
	}
}
