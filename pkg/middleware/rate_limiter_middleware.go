package middleware

import (
	"net/http"

	"github.com/1107-adishjain/sentinel/pkg/logger"
	"github.com/1107-adishjain/sentinel/pkg/metrics"
	"github.com/1107-adishjain/sentinel/pkg/models"
	"github.com/1107-adishjain/sentinel/pkg/ratelimiter"
	"github.com/gin-gonic/gin"
)

func RateLimiterMiddleware(limiter ratelimiter.Limiter, log *logger.RateLimitLogger) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, exists := c.Get(ContextUserIDKey)
		if !exists {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		metrics.TotalRequests.Inc()
		allowed, err := limiter.Allow(userID.(string))
		if err != nil || !allowed {

			metrics.BlockedRequests.Inc()
			event := models.RateLimitEvent{
				UserID:   userID.(string),
				Endpoint: c.FullPath(),
				Method:   c.Request.Method,
				IP:       c.ClientIP(),
				Reason:   "rate_limit_exceeded",
			}

			// Non-blocking log signal
			log.Log(event)

			c.AbortWithStatus(http.StatusTooManyRequests)
			return
		}
		metrics.AllowedRequests.Inc()

		c.Next()
	}
}
