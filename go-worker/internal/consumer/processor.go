package consumer

import (
	"context"
	"encoding/json"
	"go-worker/internal/db"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ChargePayload struct {
	LogID       string `json:"log_id"`
	Name        string `json:"name"`
	Amount      int    `json:"amount"`
	ReferenceID string `json:"reference_id"`
}

func ProcessChargeMessage(data []byte) {
	var payload ChargePayload

	if err := json.Unmarshal(data, &payload); err != nil {
		log.Printf("Invalid message format: %v", err)
		return
	}
	// fmt.Println("Processing charge message:", payload)
	// Create a context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	mongoURI := os.Getenv("MONGO_URI")
	if mongoURI == "" {
		mongoURI = "mongodb://localhost:27017"
	}

	dbName := os.Getenv("MONGO_DB_NAME")
	if dbName == "" {
		dbName = "donation_db"
	}

	mongoLogger, err := db.NewMongoLogger(ctx, mongoURI, dbName, "charge_processing_logs")
	if err != nil {
		log.Printf("Failed to create MongoDB logger: %v", err)
		return
	}

	logProducer, err := db.NewMongoLogger(ctx, mongoURI, "go-api", "charge_logs")
	if err != nil {
		log.Printf("Failed to create MongoDB logger for charge logs: %v", err)
		return
	}

	logEntry, err := logProducer.GetLogByID(ctx, payload.LogID)
	if err != nil {
		log.Printf("Failed to retrieve log entry: %v", err)
		return
	}

	insertedID, err := mongoLogger.LogChargeProcess(ctx, payload.LogID, "processing", payload.Amount, "Processing donation")
	if err != nil {
		log.Printf("Failed to log charge process: %v", err)
		return
	}

	if objectID, ok := insertedID.(primitive.ObjectID); ok {
		payload.LogID = objectID.Hex()
	} else {
		log.Printf("Failed to convert insertedID to ObjectID")
		return
	}

	// Simulate a delay
	time.Sleep(3 * time.Second)
	// Log the processing details
	log.Printf("Processing donation for %s, amount: %d, log entry: %v, log ID: %s", payload.Name, payload.Amount, logEntry, payload.LogID)
}
