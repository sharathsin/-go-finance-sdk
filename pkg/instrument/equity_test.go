package instrument

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEquity(t *testing.T) {
	eq := NewEquity("AAPL-US", "USD", "AAPL")

	assert.Equal(t, "AAPL-US", eq.ID())
	assert.Equal(t, "USD", eq.Currency())
	assert.Equal(t, "AAPL", eq.Symbol())
	assert.Equal(t, TypeEquity, eq.Type())
}
