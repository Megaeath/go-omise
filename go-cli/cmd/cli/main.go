package main

import (
	"bytes"
	"fmt"
	"go-cli/internal/chargeclient" // Corrected import path
	"go-cli/internal/csvreader"    // Corrected import path
	"go-cli/internal/model"        // Corrected import path
	"go-cli/internal/summary"      // Corrected import path
	"io/ioutil"
	"os"
	"sync"

	"github.com/schollz/progressbar/v3"
)

const DefaultConcurrency = 5

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
	results := processDonations(rows, DefaultConcurrency)

	summary.PrintSummary(results)
}

func processDonations(rows []model.DonationRow, concurrency int) []model.DonationResult {
	if concurrency <= 0 {
		concurrency = DefaultConcurrency
	}

	results := make([]model.DonationResult, len(rows))
	bar := progressbar.New(len(rows))

	jobs := make(chan int, len(rows))
	var wg sync.WaitGroup

	for w := 0; w < concurrency; w++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := range jobs {
				row := rows[i]
				ok, err := chargeclient.SendCharge(row)

				results[i] = model.DonationResult{
					Row:     row,
					Success: ok,
					Error:   err,
				}
				bar.Add(1)
			}
		}()
	}

	for i := range rows {
		jobs <- i
	}
	close(jobs)

	wg.Wait()
	return results
}
