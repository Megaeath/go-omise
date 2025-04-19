package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"

	"go-cli/chargeclient"
	"go-cli/csvreader"
	"go-cli/decrypt"
	"go-cli/model"
	"go-cli/summary"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: cli-donation <path-to-file>")
		os.Exit(1)
	}

	filePath := os.Args[1]
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		panic(err)
	}

	// Detect encrypted (simple check for non-ASCII)
	if bytes.IndexFunc(data, func(r rune) bool { return r > 127 }) != -1 {
		data = decrypt.ROT128(data)
	}

	rows, err := csvreader.ParseCSV(bytes.NewReader(data))
	if err != nil {
		fmt.Println("Error parsing CSV:", err)
		panic(err)
	}

	fmt.Println("performing donations...")
	var results []model.DonationResult
	for _, row := range rows {
		ok, err := chargeclient.SendMockCharge(row)
		results = append(results, model.DonationResult{
			Row:     row,
			Success: ok,
			Error:   err,
		})
	}
	fmt.Println("done.")

	summary.PrintSummary(results)
}
