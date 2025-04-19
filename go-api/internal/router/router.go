package router

import (
	"go-api/internal/handler"
	"go-api/internal/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	api := r.Group("/api", middleware.RateLimiter())
	{
		api.POST("/charge", handler.ChargeHandler)
	}

	return r
}
