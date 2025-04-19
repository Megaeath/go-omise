package model

import "time"

type ChargeRequestLog struct {
	ID           string    `bson:"_id,omitempty"`
	Name         string    `bson:"name"`
	Amount       int       `bson:"amount_subunits"`
	Timestamp    time.Time `bson:"timestamp"`
	Status       string    `bson:"status"` // e.g., "queued", "failed_validation"
	ErrorMessage string    `bson:"error_message,omitempty"`
	MaskedCard   string    `bson:"masked_card"` // e.g., "************1234"
	ReferenceID  string    `bson:"reference_id"`
}
