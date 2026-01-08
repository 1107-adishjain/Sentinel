package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/1107-adishjain/sentinel/pkg/ratelimiter"
)

func RateLimiterMiddleware(limiter ratelimiter.Limiter) gin.HandlerFunc {
	return func(c *gin.Context) {

		userID, exists := c.Get(ContextUserIDKey)
		if !exists {
			c.AbortWithStatusJSON(401, gin.H{
				"error": "Unauthorized",
			})
			return
		}

		allowed,err:= limiter.Allow(userID.(string))
		if err != nil || !allowed {
			c.AbortWithStatus(http.StatusTooManyRequests)
			return
		}
		c.Next()
	}
}
