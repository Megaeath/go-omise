package consumer

import (
	"encoding/json"
	"testing"
)

func TestProcessChargeMessage_ValidPayload(t *testing.T) {
	payload := ChargePayload{
		LogID:       "test-log-id",
		Name:        "Test User",
		Amount:      100,
		ReferenceID: "test-ref-id",
	}
	data, _ := json.Marshal(payload)

	ProcessChargeMessage(data)
	// Add assertions or mock checks if necessary
}

func TestProcessChargeMessage_InvalidPayload(t *testing.T) {
	data := []byte("invalid-json")

	ProcessChargeMessage(data)
	// Add assertions or mock checks if necessary
}
