package routes

import (
	"github.com/1107-adishjain/sentinel/app"
	cont "github.com/1107-adishjain/sentinel/pkg/controllers"
	mw "github.com/1107-adishjain/sentinel/pkg/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/1107-adishjain/sentinel/pkg/metrics"
	"github.com/prometheus/client_golang/prometheus/promhttp"

)

func Routes(app *app.Application) *gin.Engine{
	router:= gin.Default()

	router.Use(mw.SecurityHeaders())
	router.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"*"},
		AllowHeaders:  []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		AllowMethods:  []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		ExposeHeaders: []string{"Content-Length"},
		MaxAge:        3600,
	}))
	metrics.Register()
	r:=router.Group("/api/v1")
	{
		r.POST("/signup",cont.Signup(app.DB))
		r.POST("/login",cont.Login(app.DB))
		r.Use(mw.AuthMiddleware())
		r.Use(mw.RateLimiterMiddleware(app.Ratelimiter,app.Logger))
		r.GET("/healthcheck",cont.Healthcheck())
		r.GET("/ping",cont.Ping())
	}
	router.GET("/metrics", gin.WrapH(promhttp.Handler()))
	return router
}