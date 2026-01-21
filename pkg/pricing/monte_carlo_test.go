package pricing

import (
	"context"
	"testing"
	"time"

	"github.com/antigravity/go-finance-sdk/pkg/instrument"
	"github.com/shopspring/decimal"
)

func TestMonteCarloOptionPricing(t *testing.T) {
	// Setup
	underlying := instrument.NewEquity("AAPL", "USD", "AAPL")
	strike := decimal.NewFromInt(100) // ATM
	expiry := time.Now().Add(30 * 24 * time.Hour)
	opt := instrument.NewEuropeanOption("OPT1", underlying, strike, expiry, instrument.Call)

	// Pricing
	// Rate=5%, Volatility=20%, Sims=10000
	pricer := NewMonteCarloPricer(10000, 0.05, 0.20)
	price, err := pricer.Price(context.Background(), opt)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if price <= 0 {
		t.Errorf("Expected positive price, got %f", price)
	}
}

func BenchmarkMonteCarloPricing(b *testing.B) {
	underlying := instrument.NewEquity("AAPL", "USD", "AAPL")
	strike := decimal.NewFromInt(150)
	expiry := time.Now().Add(30 * 24 * time.Hour)
	opt := instrument.NewEuropeanOption("OPT1", underlying, strike, expiry, instrument.Call)
	pricer := NewMonteCarloPricer(10000, 0.05, 0.20)
	ctx := context.Background()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = pricer.Price(ctx, opt)
	}
}
