package ratelimiter

// rateLimiterLua contains the Redis Lua script.
// Kept in Go to allow preloading and atomic execution.
const rateLimiterLua = `
-- Token Bucket Rate Limiter (Atomic)
local key = KEYS[1]
local max_tokens = tonumber(ARGV[1])
local refill_rate = tonumber(ARGV[2])
local now = tonumber(ARGV[3])
local ttl = tonumber(ARGV[4])

local data = redis.call("HMGET", key, "tokens", "last_refill")
local tokens = tonumber(data[1]) or max_tokens
local last_refill = tonumber(data[2]) or now

local elapsed = math.max(0, now - last_refill)
local refill = elapsed * refill_rate
tokens = math.min(max_tokens, tokens + refill)

local allowed = 0
if tokens >= 1 then
	tokens = tokens - 1
	allowed = 1
end

redis.call("HMSET", key, "tokens", tokens, "last_refill", now)
redis.call("EXPIRE", key, ttl)

return allowed
`
