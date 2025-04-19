package handler

import (
	"go-api/internal/model"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/google/uuid"
)

func ChargeHandler(c *gin.Context) {
	var req model.ChargeRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request: " + err.Error()})
		return
	}

	// 🪪 Mask card
	// maskedCard := "XXXX-XXXX-XXXX-" + req.CCNumber[len(req.CCNumber)-4:]

	// 🧾 Generate log ID
	logID := uuid.NewString()

	// 📝 Simulate logging to DB and sending to Kafka (placeholder)
	// logToDB(logID, req.Name, maskedCard, req.AmountSubunits)
	// sendToKafka(logID, maskedCard, req.AmountSubunits)

	c.JSON(http.StatusOK, model.ChargeResponse{LogID: logID})
}
