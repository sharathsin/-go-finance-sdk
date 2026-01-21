package market

import (
	"context"
	"errors"
	"math/rand"
	"time"

	"github.com/antigravity/go-finance-sdk/pkg/common/observability"
	"github.com/antigravity/go-finance-sdk/pkg/common/resilience"
	"github.com/shopspring/decimal"
	"go.uber.org/zap"
	"golang.org/x/time/rate"
)

// MockClient simulates an external API client with built-in resilience.
type MockClient struct {
	breaker *resilience.CircuitBreaker
	limiter *resilience.RateLimiter
	logger  *zap.Logger
}

// NewMockClient creates a new mock client.
func NewMockClient() *MockClient {
	return &MockClient{
		breaker: resilience.NewCircuitBreaker("market-api", 5*time.Second),
		limiter: resilience.NewRateLimiter(rate.Limit(10), 1), // 10 req/sec
		logger:  observability.GetLogger(),
	}
}

// GetPrice fetches the price using resilience patterns.
func (c *MockClient) GetPrice(ctx context.Context, symbol string) (Price, error) {
	logger := c.logger.With(zap.String("symbol", symbol))

	// 1. Rate Limiting
	if err := c.limiter.Wait(ctx); err != nil {
		logger.Error("Rate limit exceeded", zap.Error(err))
		return Price{}, err
	}

	var price Price

	// 2. Retry Logic
	op := func(ctx context.Context) error {
		// 3. Circuit Breaker
		_, err := c.breaker.Execute(func() (interface{}, error) {
			start := time.Now()
			p, err := c.simulateNetworkCall(ctx, symbol)
			duration := time.Since(start).Seconds()

			// Observability: Metrics
			status := "success"
			if err != nil {
				status = "error"
				observability.ExternalApiErrors.WithLabelValues("mock", "network").Inc()
			}
			observability.HttpRequestDuration.WithLabelValues("mock", "GetPrice", status).Observe(duration)

			if err != nil {
				return nil, err
			}
			price = p
			return p, nil
		})
		return err
	}

	err := resilience.Retry(ctx, resilience.DefaultBackoffConfig, op)
	if err != nil {
		logger.Error("Failed to fetch price after retries", zap.Error(err))
		return Price{}, err
	}

	logger.Info("Successfully fetched price", zap.String("price", price.Value.String()))
	return price, nil
}

// simulateNetworkCall simulates a flaky network call.
func (c *MockClient) simulateNetworkCall(ctx context.Context, symbol string) (Price, error) {
	// Simulate latency
	select {
	case <-time.After(time.Duration(rand.Intn(100)) * time.Millisecond):
	case <-ctx.Done():
		return Price{}, ctx.Err()
	}

	// Simulate random failure (20% chance)
	if rand.Float64() < 0.2 {
		return Price{}, errors.New("simulated network error 500")
	}

	// Generate random price
	val := 100.0 + rand.Float64()*50.0
	return Price{
		Symbol:    symbol,
		Value:     decimal.NewFromFloat(val).Round(2),
		Timestamp: time.Now(),
	}, nil
}
