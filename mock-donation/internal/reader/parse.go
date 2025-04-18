package reader

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"

	"mockdonate/internal/model"
)

func ReadDonationsStream(path string) ([]model.Donation, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	reader := csv.NewReader(f)
	var donations []model.Donation

	// Skip header
	_, err = reader.Read()
	if err != nil {
		return nil, fmt.Errorf("failed to read header: %w", err)
	}

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("failed to read record: %w", err)
		}

		// Parse fields
		amount, err := strconv.Atoi(record[1])
		if err != nil {
			return nil, fmt.Errorf("invalid amount in record: %v, error: %w", record, err)
		}

		expMonth, err := strconv.Atoi(record[4])
		if err != nil {
			return nil, fmt.Errorf("invalid expiration month in record: %v, error: %w", record, err)
		}

		expYear, err := strconv.Atoi(record[5])
		if err != nil {
			return nil, fmt.Errorf("invalid expiration year in record: %v, error: %w", record, err)
		}

		// Create Donation struct
		d := model.Donation{
			Name:           record[0],
			AmountSubunits: amount,
			CCNumber:       record[2],
			CVV:            record[3],
			ExpMonth:       expMonth,
			ExpYear:        expYear,
		}

		donations = append(donations, d)
	}

	return donations, nil
}
