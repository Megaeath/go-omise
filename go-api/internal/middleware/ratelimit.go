package middleware

import (
	"fmt"
	"go-api/internal/config"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	rateLimitWindow = 3600           // seconds
	maxRequests     = 10000000       // max per window
	globalKey       = "global_limit" // key for all users
)

func RateLimiter() gin.HandlerFunc {
	return func(c *gin.Context) {
		count, err := config.RedisClient.Incr(config.Ctx, globalKey).Result()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Rate limiter error"})
			c.Abort()
			return
		}

		if count == 1 {
			// set TTL if it's a new key
			config.RedisClient.Expire(config.Ctx, globalKey, time.Second*rateLimitWindow)
		}

		if count > maxRequests {
			ttl, _ := config.RedisClient.TTL(config.Ctx, globalKey).Result()
			c.JSON(http.StatusTooManyRequests, gin.H{
				"error":        "Rate limit exceeded",
				"retry_after":  fmt.Sprintf("%.0f seconds", ttl.Seconds()),
				"current_hits": count,
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
