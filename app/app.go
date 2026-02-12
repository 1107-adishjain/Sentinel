package app

import (
	"github.com/1107-adishjain/sentinel/pkg/config"
	"github.com/1107-adishjain/sentinel/pkg/logger"
	"github.com/1107-adishjain/sentinel/pkg/ratelimiter"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type Application struct {
	Config      *config.Config
	RedisClient *redis.Client
	Ratelimiter ratelimiter.Limiter
	DB          *gorm.DB
	Logger      *logger.RateLimitLogger
}
