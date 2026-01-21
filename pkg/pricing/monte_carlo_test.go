package pricing

import (
	"testing"
	"time"

	"github.com/antigravity/go-finance-sdk/pkg/instrument"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

func TestMonteCarloPricer(t *testing.T) {
	r := 0.05
	sigma := 0.2
	sims := 1000
	pricer := NewMonteCarloPricer(sims, r, sigma)

	underlying := instrument.NewEquity("TEST", "USD", "TEST")
	expiry := time.Now().Add(365 * 24 * time.Hour)
	strike := decimal.NewFromInt(100)

	opt := instrument.NewEuropeanOption("OPT", underlying, strike, expiry, instrument.Call)

	price, err := pricer.Price(opt)
	assert.NoError(t, err)
	assert.True(t, price > 0)

	// MC should coverage to BS price with enough sims, but for 1000 it might be noisy.
	// But it should be roughly around 10.45
	assert.InDelta(t, 10.45, price, 2.0)
}
