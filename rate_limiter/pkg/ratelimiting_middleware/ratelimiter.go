package ratelimiting_middleware

//RateLimiter contract for ratelimiting provider
type RateLimiter interface {
	RateLimit() (bool)
}
