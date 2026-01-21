package resilience

import (
	"errors"
	"time"

	"github.com/sony/gobreaker"
)

// CircuitBreaker wraps gobreaker.CircuitBreaker to provide a simplified interface.
type CircuitBreaker struct {
	cb *gobreaker.CircuitBreaker
}

// NewCircuitBreaker creates a new circuit breaker with generic settings.
// name: used for logging/metrics
// timeout: how long to keep the circuit open before checking if service is back
func NewCircuitBreaker(name string, timeout time.Duration) *CircuitBreaker {
	st := gobreaker.Settings{
		Name:    name,
		Timeout: timeout,
		ReadyToTrip: func(counts gobreaker.Counts) bool {
			// Trip if > 3 failures
			// And failure ratio > 40%
			failureRatio := float64(counts.TotalFailures) / float64(counts.Requests)
			return counts.Requests >= 3 && failureRatio >= 0.4
		},
		OnStateChange: func(name string, from gobreaker.State, to gobreaker.State) {
			// In a real app, you would log this to structured logger
			// fmt.Printf("Circuit Breaker '%s' changed from %s to %s\n", name, from, to)
		},
	}
	return &CircuitBreaker{
		cb: gobreaker.NewCircuitBreaker(st),
	}
}

// Execute runs the given function through the circuit breaker.
func (c *CircuitBreaker) Execute(req func() (interface{}, error)) (interface{}, error) {
	return c.cb.Execute(req)
}

// ExecuteErr is a helper for functions that only return an error.
func (c *CircuitBreaker) ExecuteErr(req func() error) error {
	_, err := c.cb.Execute(func() (interface{}, error) {
		return nil, req()
	})
	return err
}

var ErrOpenState = errors.New("circuit breaker is open")
