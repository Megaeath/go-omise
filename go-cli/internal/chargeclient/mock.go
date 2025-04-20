package chargeclient

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"go-cli/internal/model"
	"net/http"
)

func SendCharge(row model.DonationRow, host string) (bool, error) {

	payload, err := json.Marshal(row)
	if err != nil {
		return false, err
	}

	url := fmt.Sprintf("%s/api/charge", host)
	// fmt.Println("Sending charge to:", url)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(payload))
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return false, errors.New("failed to process charge: " + resp.Status)
	}

	return true, nil
}
