package router

import (
	"go-api/internal/handler"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	api := r.Group("/api")
	{
		api.POST("/charge", handler.ChargeHandler)
	}

	return r
}
