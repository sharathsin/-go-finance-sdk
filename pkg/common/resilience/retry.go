package resilience

import (
	"context"
	"math/rand"
	"time"
)

// BackoffConfig configures the retry behavior.
type BackoffConfig struct {
	MaxRetries int
	BaseDelay  time.Duration
	MaxDelay   time.Duration
	Jitter     bool
}

// DefaultBackoffConfig provides sensible defaults.
var DefaultBackoffConfig = BackoffConfig{
	MaxRetries: 3,
	BaseDelay:  100 * time.Millisecond,
	MaxDelay:   2 * time.Second,
	Jitter:     true,
}

// Retry executes the operation op with exponential backoff.
// It returns the error of the last attempt if all retries fail.
// It respects context cancellation.
func Retry(ctx context.Context, cfg BackoffConfig, op func(ctx context.Context) error) error {
	var err error
	for i := 0; i <= cfg.MaxRetries; i++ {
		// Attempt the operation
		if err = op(ctx); err == nil {
			return nil
		}

		// If this was the last attempt, return the error
		if i == cfg.MaxRetries {
			return err
		}

		// Calculate backoff
		power := 1 << i
		backoff := float64(cfg.BaseDelay) * float64(power) // 2^i * BaseDelay
		if backoff > float64(cfg.MaxDelay) {
			backoff = float64(cfg.MaxDelay)
		}

		if cfg.Jitter {
			// Add +/- 20% jitter
			jitterFactor := 0.8 + rand.Float64()*0.4
			backoff *= jitterFactor
		}

		delay := time.Duration(backoff)

		// Wait or cancel
		select {
		case <-time.After(delay):
			continue
		case <-ctx.Done():
			return ctx.Err()
		}
	}
	return err
}
