package main

import (
	"encoding/csv"
	"fmt"
	"go-tamboon/cipher"
	"io"
	"os"
	"strconv"
)

type Donation struct {
	Name           string `csv:"Name"`
	AmountSubunits int    `csv:"AmountSubunits"`
	CCNumber       string `csv:"CCNumber"`
	CVV            string `csv:"CVV"`
	ExpMonth       int    `csv:"ExpMonth"`
	ExpYear        int    `csv:"ExpYear"`
}

func main() {
	cwd, _ := os.Getwd()
	fmt.Println("Current working directory:", cwd)

	encryptedFile := "data/fng.1000.csv.rot128"
	file, err := os.Open(encryptedFile)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	outputFile := "data/fng.csv"
	out, err := os.Create(outputFile)
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}
	defer out.Close()

	reader, err := cipher.NewRot128Reader(file)

	_, err = io.Copy(out, reader)
	if err != nil {
		fmt.Println("Error decoding file:", err)
		return
	}

	fmt.Println("Decryption successful! Output file:", outputFile)

	donations, err := readDonations(outputFile)
	if err != nil {
		fmt.Println("Error reading donations:", err)
		return
	}

	for _, d := range donations {
		fmt.Printf("%+v\n", d)
	}
}

func readDonations(csvPath string) ([]Donation, error) {
	file, err := os.Open(csvPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	// ensures the CSV has exactly 6 fields per line
	reader.FieldsPerRecord = 6

	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var donations []Donation
	for i, record := range records {
		if i == 0 {
			// skip header
			continue
		}

		amount, _ := strconv.Atoi(record[1])
		expMonth, _ := strconv.Atoi(record[4])
		expYear, _ := strconv.Atoi(record[5])

		d := Donation{
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
