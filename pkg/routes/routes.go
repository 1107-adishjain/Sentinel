package routes

import (
	cont "github.com/1107-adishjain/sentinel/pkg/controllers"
	"github.com/gin-gonic/gin"
)

func Routes() *gin.Engine{
	router:= gin.Default()


	r:=router.Group("/api/v1")
	{
		r.GET("/healthcheck",cont.Healthcheck())
		r.GET("/ping",cont.Ping())
	}

	return router
}