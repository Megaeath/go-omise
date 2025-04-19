package chargeclient

import (
	"errors"
	"math/rand"
	"time"

	"go-cli/internal/model"
)

func SendMockCharge(row model.DonationRow) (bool, error) {
	time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
	if rand.Float64() < 0.9 {
		return true, nil
	}
	return false, errors.New("mock charge failed")
}
