package ratelimiter

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

// RedisLimiter implements token bucket rate limiting using Redis + Lua.
type RedisLimiter struct {
	client       *redis.Client
	script       *redis.Script
	maxTokens    int
	refillRate   int // tokens per second
	keyTTL       time.Duration
	ctx          context.Context
}

// NewRedisLimiter constructs a Redis-backed rate limiter.
// All configuration details lives here, not in middleware.
func NewRedisLimiter(client *redis.Client) *RedisLimiter {
	return &RedisLimiter{
		client:     client,
		script:     redis.NewScript(rateLimiterLua),
		maxTokens:  10,
		refillRate: 1,
		keyTTL:     time.Hour,
		ctx:        context.Background(),
	}
}

// Allow checks whether a request for the given key is allowed.
func (r *RedisLimiter) Allow(key string) (bool, error) {
	now := time.Now().Unix()

	// lua returns 1 if allowed, 0 if not
	res, err := r.script.Run(
		r.ctx,
		r.client,
		[]string{"rate_limit:" + key},
		r.maxTokens,
		r.refillRate,
		now,
		r.keyTTL.Seconds(),
	).Int()

	if err != nil {
		return false, err
	}

	// 1 means allowed request 
	return res == 1, nil
}