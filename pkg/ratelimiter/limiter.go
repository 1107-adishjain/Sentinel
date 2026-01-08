package ratelimiter

// Limiter defines the contract for any rate limiting backend.
// Middleware does not care how limits are enforced.
type Limiter interface {
	// Allow returns whether the request identified by key is allowed.
	Allow(key string) (bool, error)
}
