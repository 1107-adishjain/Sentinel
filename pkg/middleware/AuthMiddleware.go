package middleware

import(
	"github.com/gin-gonic/gin"
)

const ContextUserIDKey = "user_id" //this is set this way so that any naming errors dont occur we have a fixed value to be used entirely everywhere

func AuthMiddleware() gin.HandlerFunc{
	return func(c *gin.Context) {
		UserID:= "user123"
		c.Set(ContextUserIDKey,UserID)
		c.Next()
	}
}