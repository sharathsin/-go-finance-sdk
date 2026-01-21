package resilience

import (
	"context"

	"golang.org/x/time/rate"
)

// RateLimiter limits the number of events per second.
type RateLimiter struct {
	limiter *rate.Limiter
}

// NewRateLimiter creates a limiter that allows events up to rate r and permits bursts of at most b tokens.
func NewRateLimiter(r rate.Limit, b int) *RateLimiter {
	return &RateLimiter{
		limiter: rate.NewLimiter(r, b),
	}
}

// Wait blocks until the limiter permits one event to happen.
func (l *RateLimiter) Wait(ctx context.Context) error {
	return l.limiter.Wait(ctx)
}
