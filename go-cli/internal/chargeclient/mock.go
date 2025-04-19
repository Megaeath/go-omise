package chargeclient

import (
	"bytes"
	"encoding/json"
	"errors"
	"go-cli/internal/model"
	"net/http"
)

func SendCharge(row model.DonationRow) (bool, error) {
	// Prepare the JSON payload
	payload, err := json.Marshal(row)
	if err != nil {
		return false, err
	}

	// Make the HTTP POST request
	resp, err := http.Post("http://localhost:8080/api/charge", "application/json", bytes.NewBuffer(payload))
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	// Check the response status code
	if resp.StatusCode != http.StatusOK {
		return false, errors.New("failed to process charge: " + resp.Status)
	}

	return true, nil
}
