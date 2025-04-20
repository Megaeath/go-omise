package main

import (
	"go-api/internal/config"
	"go-api/internal/router"
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	requestCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests",
		},
		[]string{"path"},
	)
)

func init() {
	prometheus.MustRegister(requestCount)
}

func main() {
	config.InitRedis()
	config.InitMongo()
	config.InitKafkaTopic()

	r := router.SetupRouter()
	http.Handle("/metrics", promhttp.Handler())

	log.Println("API running at :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
