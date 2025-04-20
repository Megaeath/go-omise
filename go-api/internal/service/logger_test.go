package service

import (
	"go-api/internal/model"
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestLogChargeRequest(t *testing.T) {
	mockLog := model.ChargeRequestLog{
		Name:        "John Doe",
		Amount:      1000,
		Timestamp:   time.Now(),
		Status:      "queued",
		MaskedCard:  "************1111",
		ReferenceID: "mock-ref-id",
	}

	// Mock db.ChargeLogs.InsertOne behavior
	// Replace this with actual mocking logic if needed
	result, err := LogChargeRequest(mockLog)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if result == primitive.NilObjectID {
		t.Fatal("Expected valid ObjectID, got NilObjectID")
	}
}
