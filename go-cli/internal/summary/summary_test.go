package summary

import (
	"testing"

	"go-cli/internal/model"
)

func TestPrintSummary(t *testing.T) {
	results := []model.DonationResult{
		{Row: model.DonationRow{Name: "A", AmountSubunits: 1000}, Success: true},
		{Row: model.DonationRow{Name: "B", AmountSubunits: 2000}, Success: true},
		{Row: model.DonationRow{Name: "C", AmountSubunits: 1000}, Success: false},
	}

	PrintSummary(results)
	// This is a visual test; real test would capture output or check logic separately
}
