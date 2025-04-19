package handler

import (
	"go-api/internal/model"
	"net/http"

	"github.com/gin-gonic/gin"

	"go-api/internal/service"
	"go-api/internal/utils"

	"github.com/google/uuid"
)

func ChargeHandler(c *gin.Context) {
	var req model.ChargeRequest

	if err := c.ShouldBindJSON(&req); err != nil {
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
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to log charge"})
		return
	}

	c.JSON(http.StatusOK, model.ChargeResponse{LogID: referenceID})
}
