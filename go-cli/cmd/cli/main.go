package main

import (
	"bytes"
	"fmt"
	"go-cli/internal/chargeclient" // Corrected import path
	"go-cli/internal/csvreader"    // Corrected import path
	// Corrected import path
	"go-cli/internal/model"   // Corrected import path
	"go-cli/internal/summary" // Corrected import path
	"io/ioutil"
	"os"
	"github.com/schollz/progressbar/v3"
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
	// if bytes.IndexFunc(data, func(r rune) bool { return r > 127 }) != -1 {
	// 	data = decrypt.ROT128(data)
	// 	fmt.Println("decrypted data", string(data))
	// }
	// fmt.Println("decrypted data:", string(data))

	rows, err := csvreader.ParseCSV(bytes.NewReader(data))
	if err != nil {
		fmt.Println("Error parsing CSV:", err)
		panic(err)
	}
	// fmt.Println("total rows:", len(rows))
	fmt.Println("performing donations...")
	bar := progressbar.New(len(rows))
	var results []model.DonationResult
	for _, row := range rows {
		// fmt.Println("donating to", row.Name, "for", row.AmountSubunits)
		ok, err := chargeclient.SendCharge(row)
		// if err != nil {
		// 	fmt.Println("Error sending charge:", err)
		// } else if !ok {
		// 	fmt.Println("Charge failed for:", row.Name)
		// }
		results = append(results, model.DonationResult{
			Row:     row,
			Success: ok,
			Error:   err,
		})
		bar.Add(1)
	}
	fmt.Println("done.")

	summary.PrintSummary(results)
}
