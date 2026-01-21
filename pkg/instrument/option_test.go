package instrument

import (
	"testing"
	"time"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

func TestOption(t *testing.T) {
	underlying := NewEquity("AAPL-US", "USD", "AAPL")
	strike := decimal.NewFromInt(150)
	expiry := time.Now().Add(30 * 24 * time.Hour)

	opt := NewEuropeanOption("AAPL-CALL-150", underlying, strike, expiry, Call)

	assert.Equal(t, "AAPL-CALL-150", opt.ID())
	assert.Equal(t, underlying, opt.Underlying())
	assert.Equal(t, strike, opt.Strike())
	assert.Equal(t, expiry, opt.Expiry())
	assert.Equal(t, Call, opt.OptionType())
	assert.Equal(t, European, opt.Style())
	assert.Equal(t, TypeOption, opt.Type())
	assert.Equal(t, "USD", opt.Currency())
}
