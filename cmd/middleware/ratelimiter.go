package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

// Rate limit middleware
//
//	Uses limiter := rate.NewLimiter(2, 4) internally
func RateLimiter() gin.HandlerFunc {
	limiter := rate.NewLimiter(8, 16)
	return func(c *gin.Context) {
		if limiter.Allow() {
			c.Next()
		} else {
			c.JSON(http.StatusTooManyRequests, gin.H{
				"message": "Limit exceed",
			})
			c.AbortWithStatus(http.StatusTooManyRequests)
			return
		}
	}
}
