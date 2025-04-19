package csvreader

import (
	"strings"
	"testing"
)

func TestParseCSV(t *testing.T) {
	csv := `Name,AmountSubunits,CCNumber,CVV,ExpMonth,ExpYear
Luke Skywalker,50000,1234567812345678,123,12,2025
`

	rows, err := ParseCSV(strings.NewReader(csv))
	if err != nil {
		t.Fatalf("parse failed: %v", err)
	}
	if len(rows) != 1 {
		t.Fatalf("expected 1 row, got %d", len(rows))
	}
	if rows[0].Name != "Luke Skywalker" {
		t.Errorf("expected name Luke Skywalker, got %s", rows[0].Name)
	}
}
