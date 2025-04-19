package model

type ChargeMessage struct {
	ReferenceID string `json:"reference_id"`
	LogID       string `json:"log_id"`
	Name        string `json:"name"`
	Amount      int    `json:"amount"`
}
