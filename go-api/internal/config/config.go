package config

import (
	"context"
	"log"

	"go-api/internal/db"
	"go-api/internal/kafka"

	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client
var Ctx = context.Background()

func InitRedis() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // 🔧 Replace if using Docker or external Redis
		Password: "",
		DB:       0,
	})
}

func InitMongo() {
	// MongoDB connection details
	mongoURI := "mongodb://localhost:27017" // 🔧 Replace with your MongoDB URI
	dbName := "go-api"                      // 🔧 Replace with your database name

	log.Println("Initializing MongoDB...")
	db.InitMongoDB(mongoURI, dbName)
	log.Println("MongoDB initialized successfully.")
}

func InitKafkaTopic() {
	broker := "localhost:29092"
	topic := "charge-topic"

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
