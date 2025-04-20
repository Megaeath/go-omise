package model

import (
	"testing"

	"github.com/go-playground/validator/v10"
)

func TestChargeRequestValidation(t *testing.T) {
	validate := validator.New()

	validRequest := ChargeRequest{
		Name:           "John Doe",
		AmountSubunits: 1000,
		CCNumber:       "4111111111111111",
		CVV:            "123",
		ExpMonth:       "12",
		ExpYear:        "2025",
	}

	if err := validate.Struct(validRequest); err != nil {
		t.Fatalf("Expected no validation error, got %v", err)
	}

	invalidRequest := ChargeRequest{
		Name:           "",
		AmountSubunits: -100,
		CCNumber:       "123",
		CVV:            "12",
		ExpMonth:       "0",
		ExpYear:        "20",
	}

	if err := validate.Struct(invalidRequest); err == nil {
		t.Fatal("Expected validation error, got none")
	}
}
