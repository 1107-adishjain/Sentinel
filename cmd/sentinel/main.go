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
	"github.com/1107-adishjain/sentinel/pkg/database"
	"github.com/1107-adishjain/sentinel/pkg/models"
	"github.com/1107-adishjain/sentinel/pkg/ratelimiter"
	"github.com/1107-adishjain/sentinel/pkg/redis"
	"github.com/1107-adishjain/sentinel/pkg/routes"
	"github.com/1107-adishjain/sentinel/pkg/logger"
)

func main() {

	cfg, err := cg.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading Config variables: %v", err)
	}else{
		log.Println("Config variables loaded successfully")
	}

	redisClient := redis.Connect(cfg)
	if redisClient == nil {
		log.Fatalf("Failed to connect to Redis. Exiting application.")
	}else{
		log.Println("Connected to Redis successfully")
	}

	db, err:= database.ConnectDB(cfg)
	if err!=nil{
		log.Fatalf("Error connecting to database:%v",err)
	}else{
		log.Println("Connected to Database successfully	")
	}
	
	err = models.MigrateRateLimitEvent(db)
	if err != nil {
		log.Fatalf("Error migrating database: %v", err)
	}else{
		log.Println("Database migration completed successfully")
	}
	err = models.MigrateUser(db)
	if err != nil {
		log.Fatalf("Error migrating database: %v", err)
	}else{
		log.Println("Database migration completed successfully")
	}


	limiter := ratelimiter.NewRedisLimiter(redisClient)

	ratelimitLogger:= logger.NewRateLimitLogger(10000) 
	// here you passed the event channel to the worker now both the worker and the logger are pointer to the same channel
	worker:= logger.NewWorker(db,ratelimitLogger.Events())
	ctx, worker_cancel := context.WithCancel(context.Background())
	worker.Start(ctx)

	app := &app.Application{
		RedisClient: redisClient,
		Ratelimiter: limiter,
		Config:      cfg,
		DB: db,
		Logger: ratelimitLogger,
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
		ratelimitLogger.Close() // close the logger channel
		worker_cancel()                // stop the worker
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
