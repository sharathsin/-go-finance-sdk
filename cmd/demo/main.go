package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/antigravity/go-finance-sdk/pkg/common/observability"
	"github.com/antigravity/go-finance-sdk/pkg/instrument"
	"github.com/antigravity/go-finance-sdk/pkg/market"
	"github.com/antigravity/go-finance-sdk/pkg/pricing"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/shopspring/decimal"
	"go.uber.org/zap"
)

func main() {
	// 1. Initialize Observability
	observability.InitLogger()
	logger := observability.GetLogger()
	defer logger.Sync()

	logger.Info("Starting Demo Application")

	// Start Prometheus metrics server
	go func() {
		logger.Info("Starting metrics server on :8080")
		http.Handle("/metrics", promhttp.Handler())
		http.ListenAndServe(":8080", nil)
	}()

	// 2. Market Data with Resilience
	client := market.NewMockClient()
	ctx := context.Background()

	// Simulate some market data fetches
	for i := 0; i < 5; i++ {
		logger.Info("Fetching price for AAPL", zap.Int("attempt", i+1))
		price, err := client.GetPrice(ctx, "AAPL")
		if err != nil {
			logger.Error("Failed to fetch price", zap.Error(err))
		} else {
			logger.Info("Price received", zap.String("price", price.Value.String()))
		}
		time.Sleep(200 * time.Millisecond)
	}

	// 3. Advanced Concurrency (Monte Carlo)
	logger.Info("Starting Monte Carlo Simulation")

	underlying := instrument.NewEquity("AAPL", "USD", "AAPL")
	strike := decimal.NewFromInt(100)
	expiry := time.Now().Add(30 * 24 * time.Hour)
	opt := instrument.NewEuropeanOption("OPT1", underlying, strike, expiry, instrument.Call)

	// Create a pricer with context cancellation support
	pricer := pricing.NewMonteCarloPricer(500000, 0.05, 0.20)

	// Demonstrate cancellation
	shortCtx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()

	logger.Info("Running simulation with short timeout (expect cancellation)")
	_, err := pricer.Price(shortCtx, opt)
	if err != nil {
		logger.Info("Simulation cancelled as expected", zap.Error(err))
	}

	// Run full simulation
	logger.Info("Running full simulation")
	start := time.Now()
	val, err := pricer.Price(context.Background(), opt)
	if err != nil {
		logger.Error("Simulation failed", zap.Error(err))
	} else {
		logger.Info("Simulation completed",
			zap.Float64("value", val),
			zap.Duration("duration", time.Since(start)),
		)
	}

	fmt.Println("\nDemo complete! Check metrics at http://localhost:8080/metrics (if running)")
}
