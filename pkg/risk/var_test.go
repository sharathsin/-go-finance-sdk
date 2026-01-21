package risk

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateHistoricalVaR(t *testing.T) {
	// Returns: -5%, -4%, -3%, -2%, -1%, 0%, 1%, 2%, 3%, 4%
	// Sorted: -0.05, -0.04, -0.03, ...
	returns := []float64{-0.05, -0.04, -0.03, -0.02, -0.01, 0.0, 0.01, 0.02, 0.03, 0.04}
	portfolioValue := 1000.0

	// 90% confidence level => 10% tail => index 1 (since 10 items * 0.1 = 1)
	// sorted[1] = -0.04
	// VaR = -(-0.04) * 1000 = 40

	var90 := CalculateHistoricalVaR(returns, 0.90, portfolioValue)
	assert.Equal(t, 40.0, var90)

	// 99% confidence => index 0
	// sorted[0] = -0.05
	// VaR = 50
	var99 := CalculateHistoricalVaR(returns, 0.99, portfolioValue)
	assert.Equal(t, 50.0, var99)
}
