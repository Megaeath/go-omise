package worker

import (
	"log"
	"mockdonate/internal/model"
	"mockdonate/internal/processor"
	"runtime"
	"time"
)

func ProcessDonations(donations []model.Donation) {
	numWorkers := runtime.NumCPU()
	jobs := make(chan model.Donation, numWorkers)
	done := make(chan bool)

	// Rate limiter
	limiter := time.Tick(10 * time.Millisecond) // 10/sec

	for i := 0; i < numWorkers; i++ {
		go func(id int) {
			for d := range jobs {
				<-limiter // rate limit
				if err := processor.MockCharge(d); err != nil {
					log.Printf("❌ Worker %d: error charging %s: %v\n", id, d.Name, err)
				}
			}
			done <- true
		}(i)
	}

	for _, d := range donations {
		jobs <- d
	}
	close(jobs)

	// wait for all workers
	for i := 0; i < numWorkers; i++ {
		<-done
	}
}
