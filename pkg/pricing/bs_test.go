package pricing

import (
	"testing"
	"time"

	"github.com/antigravity/go-finance-sdk/pkg/instrument"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

func TestBlackScholesPricer_Call(t *testing.T) {
	// S=100 (hardcoded in bs.go), K=100, T=1 year, r=0.05, sigma=0.2
	// Expected Call Price ~ 10.4506

	r := 0.05
	sigma := 0.2
	pricer := NewBlackScholesPricer(r, sigma)

	underlying := instrument.NewEquity("TEST", "USD", "TEST")
	expiry := time.Now().Add(365 * 24 * time.Hour)
	strike := decimal.NewFromInt(100)

	opt := instrument.NewEuropeanOption("OPT", underlying, strike, expiry, instrument.Call)

	price, err := pricer.Price(opt)
	assert.NoError(t, err)
	assert.InDelta(t, 10.45, price, 0.1)
}

func TestBlackScholesPricer_Put(t *testing.T) {
	// S=100, K=100, T=1, r=0.05, sigma=0.2
	// Expected Put Price ~ 5.5735

	r := 0.05
	sigma := 0.2
	pricer := NewBlackScholesPricer(r, sigma)

	underlying := instrument.NewEquity("TEST", "USD", "TEST")
	expiry := time.Now().Add(365 * 24 * time.Hour)
	strike := decimal.NewFromInt(100)

	opt := instrument.NewEuropeanOption("OPT", underlying, strike, expiry, instrument.Put)

	price, err := pricer.Price(opt)
	assert.NoError(t, err)
	assert.InDelta(t, 5.57, price, 0.1)
}
