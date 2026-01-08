package app

import (
	"github.com/1107-adishjain/sentinel/pkg/ratelimiter"
	"github.com/redis/go-redis/v9"
	"github.com/1107-adishjain/sentinel/pkg/config"
)
type Application struct {
	Config *config.Config
	RedisClient *redis.Client
	Ratelimiter ratelimiter.Limiter
}