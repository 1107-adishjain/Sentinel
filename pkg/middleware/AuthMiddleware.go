package middleware

import(
	"github.com/gin-gonic/gin"
)

const ContextUserIDKey = "user_id"

func AuthMiddleware() gin.HandlerFunc{
	return func(c *gin.Context) {
		UserID:= "user123"

		c.Set(ContextUserIDKey,UserID)
		c.Next()
	}
}