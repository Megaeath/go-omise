package csvreader

import (
	"encoding/csv"
	"io"
	"strconv"

	"go-cli/internal/model"
)

func ParseCSV(r io.Reader) ([]model.DonationRow, error) {
	reader := csv.NewReader(r)
	reader.TrimLeadingSpace = true

	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var rows []model.DonationRow
	for i, rec := range records {
		if i == 0 {
			continue // skip header
		}
		amt, _ := strconv.ParseInt(rec[1], 10, 64)
		rows = append(rows, model.DonationRow{
			Name:           rec[0],
			AmountSubunits: amt,
			CCNumber:       rec[2],
			CVV:            rec[3],
			ExpMonth:       rec[4],
			ExpYear:        rec[5],
		})
	}
	return rows, nil
}
