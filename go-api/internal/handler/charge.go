package handler

import (
	"fmt"
	"go-api/internal/model"
	"net/http"

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

	if err := service.LogChargeRequest(logEntry); err != nil {
		fmt.Println("Failed to log charge:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to log charge"})
		return
	}

	producer := kafka.NewProducer("localhost:9092", "charge-topic") // Replace with your broker and topic
	defer producer.Close()

	msg := model.ChargeMessage{ReferenceID: referenceID}
	if err := producer.SendChargeMessage(msg); err != nil {
		fmt.Println("Failed to send message to Kafka:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to queue charge"})
		return
	}

	c.JSON(http.StatusOK, model.ChargeResponse{LogID: referenceID})
}
