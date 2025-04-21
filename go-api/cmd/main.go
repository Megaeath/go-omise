package main

import (
	"go-api/internal/config"
	"go-api/internal/router"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	requestCount = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests",
		},
	)
)

func init() {
	prometheus.MustRegister(requestCount)
}
func handler(w http.ResponseWriter, r *http.Request) {
	requestCount.Inc()
	w.Write([]byte("Hello, Prometheus!"))
}

func main() {
	config.LoadEnv()
	config.InitRedis()
	config.InitMongo()
	config.InitKafkaTopic()

	r := router.SetupRouter()
	r.GET("/metrics", gin.WrapH(promhttp.Handler()))
	log.Println("API running at :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
