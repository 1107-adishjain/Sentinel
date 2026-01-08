package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/1107-adishjain/sentinel/app"
	cg "github.com/1107-adishjain/sentinel/pkg/config"
	"github.com/1107-adishjain/sentinel/pkg/ratelimiter"
	"github.com/1107-adishjain/sentinel/pkg/redis"
	"github.com/1107-adishjain/sentinel/pkg/routes"
)

func main() {

	cfg, err := cg.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading Config variables: %v", err)
	}

	redisClient := redis.Connect(cfg)
	if redisClient == nil {
		log.Fatalf("Failed to connect to Redis. Exiting application.")
	}

	limiter := ratelimiter.NewRedisLimiter(redisClient)

	app := &app.Application{
		RedisClient: redisClient,
		Ratelimiter: limiter,
		Config:      cfg,
	}
	srv := &http.Server{
		Addr:         ":" + cfg.PORT,
		Handler:      routes.Routes(app),
		IdleTimeout:  60 * time.Minute,
		ReadTimeout:  10 * time.Minute,
		WriteTimeout: 10 * time.Minute,
	}

	Shutdown := make(chan error)
	go func() {
		// this go routine would be listeing for the shutdown signal
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
		s := <-quit
		log.Printf("Recieved the shutdown signal:%s", s)
		ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
		defer cancel()
		Shutdown <- srv.Shutdown(ctx)
	}()

	log.Printf("Starting the server on port %s", cfg.PORT)
	err = srv.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		log.Fatalf("Error starting the server: %s", err)
	}

	shutdownErr := <-Shutdown
	if shutdownErr != nil {
		log.Fatalf("Error during server shutdown: %s", shutdownErr)
	}
	log.Println("Server shutdown gracefully")

}
