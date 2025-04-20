package kafka

import (
	"go-api/internal/model"
	"testing"
)

func TestSendChargeMessage(t *testing.T) {
	producer := NewProducer("mock-broker:29092", "mock-topic")
	defer producer.Close()

	msg := model.ChargeMessage{
		ReferenceID: "mock-ref-id",
		LogID:       "mock-log-id",
		Name:        "John Doe",
		Amount:      1000,
	}

	err := producer.SendChargeMessage(msg)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
}
