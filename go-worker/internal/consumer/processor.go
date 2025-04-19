package consumer

import (
	"context"
	"encoding/json"
	"go-worker/internal/db"
	"log"
	"time"
)

type ChargePayload struct {
	LogID       string `json:"log_id"`
	Name        string `json:"name"`
	Amount      int64  `json:"amount"`
	ReferenceID string `json:"reference_id"`
}

func ProcessChargeMessage(data []byte) {
	var payload ChargePayload

	// Unmarshal the incoming message
	if err := json.Unmarshal(data, &payload); err != nil {
		log.Printf("Invalid message format: %v", err)
		return
	}

	// Create a context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Initialize MongoDB logger for charge processing logs
	mongoLogger, err := db.NewMongoLogger(ctx, "mongodb://localhost:27017", "donation_db", "charge_processing_logs")
	if err != nil {
		log.Printf("Failed to create MongoDB logger: %v", err)
		return
	}

	// Retrieve the log entry from the charge logs collection
	logProducer, err := db.NewMongoLogger(ctx, "mongodb://localhost:27017", "go-api", "charge_logs")
	if err != nil {
		log.Printf("Failed to create MongoDB logger for charge logs: %v", err)
		return
	}

	logEntry, err := logProducer.GetLogByID(ctx, payload.ReferenceID)
	if err != nil {
		log.Printf("Failed to retrieve log entry: %v", err)
		return
	}

	// Log the charge processing status
	err = mongoLogger.LogChargeProcess(ctx, payload.LogID, "processing", payload.Amount, "Processing donation")
	if err != nil {
		log.Printf("Failed to log charge process: %v", err)
		return
	}

	// Log the processing details
	log.Printf("Processing donation for %s, amount: %d, log entry: %v", payload.Name, payload.Amount, logEntry)
}
