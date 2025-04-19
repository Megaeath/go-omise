package consumer

import (
	"encoding/json"
	"log"
)

type ChargePayload struct {
	LogID     string `json:"log_id"`
	Name      string `json:"name"`
	Amount    int64  `json:"amount"`
	Reference string `json:"reference"`
}

func ProcessChargeMessage(data []byte) {
	var payload ChargePayload

	if err := json.Unmarshal(data, &payload); err != nil {
		log.Printf("invalid message format: %v", err)
		return
	}

	// TODO: Add logic to call mock API / simulate charge
	log.Printf("Processing donation for %s, amount: %d", payload.Name, payload.Amount)

	// TODO: Log result to MongoDB with reference to LogID
}
