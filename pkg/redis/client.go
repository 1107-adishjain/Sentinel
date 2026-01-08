package redis

import (
	"context"
	"log"
	"time"
	"github.com/1107-adishjain/sentinel/pkg/config"
	"github.com/redis/go-redis/v9"
)



func Connect(cfg *config.Config) *redis.Client{
	client := redis.NewClient(&redis.Options{
		Addr:         cfg.RedisAddr,
		Password:     cfg.RedisPassword,
		DB:           0,
		DialTimeout:  5 * time.Second,
		ReadTimeout:  3 * time.Second,
		WriteTimeout: 3 * time.Second,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	if err := client.Ping(ctx).Err(); err != nil {
		log.Printf("Failed to connect to Redis: %v", err)
		return nil
	}
	return client
}
