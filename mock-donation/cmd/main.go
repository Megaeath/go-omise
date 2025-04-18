package main

import (
	"log"
	"mockdonate/internal/reader"
	"mockdonate/internal/worker"
)

func main() {
	donations, err := reader.ReadDonationsStream("fng.csv")
	if err != nil {
		log.Fatal("Failed to read donations:", err)
	}

	worker.ProcessDonations(donations)
}
