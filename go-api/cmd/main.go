package main

import (
	"go-api/internal/config"
	"go-api/internal/router"
	"log"
)

func main() {
	config.LoadEnv()
	config.InitRedis()
	config.InitMongo()
	config.InitKafkaTopic()

	r := router.SetupRouter()

	log.Println("API running at :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
