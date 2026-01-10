package ratelimiter

// This interface can be now used to define different rate limiter implementations that is any struct or any type can implement this method.  the way they want to implement rate limiting
type Limiter interface {
	// Allow returns whether the request identified by key is allowed.
	Allow(key string) (bool, error)
}
