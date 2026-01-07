package middleware

import (
	"github.com/gin-gonic/gin"
)

func RateLimiterMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		userID, exists := c.Get(ContextUserIDKey)
		if !exists {
			c.AbortWithStatusJSON(401, gin.H{
				"error": "Unauthorized",
			})
			return
		}

		// userID is now available for rate limiting logic
		_ = userID // placeholder, to be used next

		// Rate limiting logic will go here (Redis later)

		c.Next()
	}
}
