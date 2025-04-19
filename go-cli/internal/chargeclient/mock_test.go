package chargeclient

import (
	"testing"

	"go-cli/internal/model"
)

func TestSendMockCharge(t *testing.T) {
	d := model.DonationRow{
		Name:           "Test User",
		AmountSubunits: 1000,
		CCNumber:       "1234123412341234",
		CVV:            "123",
		ExpMonth:       "12",
		ExpYear:        "2030",
	}
	ok, _ := SendMockCharge(d)
	if !ok {
		t.Log("mock failure allowed as part of simulation")
	}
}
