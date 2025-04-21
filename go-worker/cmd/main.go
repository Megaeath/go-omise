package main

import (
	"go-worker/internal/consumer"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	// Load .env.local or .env.docker
	err := LoadEnv()
	if err != nil {
		log.Fatalf("failed to load env: %v", err)
	}

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
	err = consumer.StartKafkaConsumer(broker, topic, group, workers)
	if err != nil {
		log.Fatalf("worker failed: %v", err)
	}
}

func LoadEnv() error {
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "local"
	}

	var envFile string
	if env == "docker" {
		envFile = ".env.docker"
	} else {
		envFile = ".env.local"
	}

	err := godotenv.Load(envFile)
	if err != nil {
		log.Printf("No %s file found or failed to load: %v", envFile, err)
		return err
	}
	return nil
}
