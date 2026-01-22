package middleware

import (
	"strings"
	"github.com/1107-adishjain/sentinel/pkg/helper"
	"github.com/gin-gonic/gin"
)

const ContextUserIDKey = "user_id" //this is set this way so that any naming errors dont occur we have a fixed value to be used entirely everywhere

func AuthMiddleware() gin.HandlerFunc{
	return func(c *gin.Context) {
		authHeader:= c.GetHeader("Authorization")
		if authHeader==""{
			c.AbortWithStatusJSON(401,gin.H{"error":"Authorization header missing"})
			return
		}
		tokenstring := strings.TrimPrefix(authHeader,"Bearer ")
		claims, err:= helper.VerifyJWT(tokenstring)
		if err!=nil{
			c.AbortWithStatusJSON(401,gin.H{"error":"Invalid or expired token"})
			return
		}
		c.Set(ContextUserIDKey,claims.UserID)
		c.Next()
	}
}