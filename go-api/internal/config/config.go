package config

import (
	"context"
	"log"
	"os"

	"go-api/internal/db"
	"go-api/internal/kafka"

	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client
var Ctx = context.Background()

func InitRedis() {
	addr := os.Getenv("REDIS_HOST")
	if addr == "" {
		addr = "localhost:6379"
	}

	RedisClient = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "",
		DB:       0,
	})
}

func InitMongo() {
	mongoURI := os.Getenv("MONGO_URI")
	if mongoURI == "" {
		mongoURI = "mongodb://localhost:27017"
	}

	dbName := os.Getenv("MONGO_DB_NAME")
	if dbName == "" {
		dbName = "go-api"
	}

	log.Println("Initializing MongoDB...")
	db.InitMongoDB(mongoURI, dbName)
	log.Println("MongoDB initialized successfully.")
}

func InitKafkaTopic() {
	broker := os.Getenv("KAFKA_BROKER")
	if broker == "" {
		broker = "localhost:29092"
	}

	topic := os.Getenv("KAFKA_TOPIC")
	if topic == "" {
		topic = "charge-topic"
	}

	log.Println("Initializing Kafka producer...")
	producer := kafka.NewProducer(broker, topic)
	defer producer.Close()

	log.Println("Ensuring Kafka topic...")
	err := kafka.EnsureTopic(broker, topic, 1, 1)
	if err != nil {
		log.Fatalf("Failed to ensure topic: %v", err)
	}
	log.Println("Kafka topic ensured successfully.")
}
