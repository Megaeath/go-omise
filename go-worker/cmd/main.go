package main

import (
	"go-worker/internal/consumer"
	"log"
	"os"
)

func main() {
	broker := os.Getenv("KAFKA_BROKER")
	if broker == "" {
		broker = "localhost:29092"
	}

	topic := os.Getenv("KAFKA_TOPIC")
	if topic == "" {
		topic = "charge-topic"
	}

	group := os.Getenv("KAFKA_GROUP")
	if group == "" {
		group = "charge-consumer"
	}

	workers := 10
	log.Println("Starting charge worker...")
	err := consumer.StartKafkaConsumer(broker, topic, group, workers)
	if err != nil {
		log.Fatalf("worker failed: %v", err)
	}
}
