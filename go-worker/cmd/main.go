// cmd/main.go
package main

import (
	"go-worker/internal/consumer"
	"log"
)

func main() {
	broker := "localhost:29092"
	topic := "charge-topic"
	group := "charge-consumer"
	workers := 5
	log.Println("Starting charge worker...")
	err := consumer.StartKafkaConsumer(broker, topic, group, workers)
	if err != nil {
		log.Fatalf("worker failed: %v", err)
	}
}
