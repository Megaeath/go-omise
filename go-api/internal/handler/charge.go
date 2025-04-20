package handler

import (
	"fmt"
	"go-api/internal/model"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"

	"go-api/internal/kafka"
	"go-api/internal/service"
	"go-api/internal/utils"

	"github.com/google/uuid"
)

func ChargeHandler(c *gin.Context) {
	var req model.ChargeRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		fmt.Println("Failed to bind JSON:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request: " + err.Error()})
		return
	}

	referenceID := uuid.New().String()
	logEntry := model.ChargeRequestLog{
		Name:        req.Name,
		Amount:      req.AmountSubunits,
		Status:      "queued",
		MaskedCard:  utils.MaskCardNumber(req.CCNumber),
		ReferenceID: referenceID,
	}

	logEntryID, err := service.LogChargeRequest(logEntry)
	if err != nil {
		fmt.Println("Failed to log charge:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to log charge"})
		return
	}

	broker := os.Getenv("KAFKA_BROKER")
	if broker == "" {
		broker = "localhost:29092"
	}

	topic := os.Getenv("KAFKA_TOPIC")
	if topic == "" {
		topic = "charge-topic"
	}

	producer := kafka.NewProducer(broker, topic)
	defer producer.Close()

	msg := model.ChargeMessage{
		LogID:  logEntryID.Hex(), // Convert ObjectID to string
		Name:   req.Name,
		Amount: req.AmountSubunits,
	}
	if err := producer.SendChargeMessage(msg); err != nil {
		fmt.Println("Failed to send message to Kafka:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to queue charge"})
		return
	}

	c.JSON(http.StatusOK, model.ChargeResponse{LogID: referenceID})
}
