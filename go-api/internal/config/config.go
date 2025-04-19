package config

import (
	"context"
	"log"

	"go-api/internal/db"

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
