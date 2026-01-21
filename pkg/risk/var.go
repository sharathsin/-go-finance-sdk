package risk

import (
	"sort"
)

// CalculateHistoricalVaR calculates Value at Risk using historical simulation.
// returns: The estimated loss amount at the given confidence level.
func CalculateHistoricalVaR(returns []float64, confidenceLevel float64, portfolioValue float64) float64 {
	if len(returns) == 0 {
		return 0.0
	}

	sortedReturns := make([]float64, len(returns))
	copy(sortedReturns, returns)
	sort.Float64s(sortedReturns)

	index := int(float64(len(sortedReturns))*(1.0-confidenceLevel) + 1e-9)
	if index < 0 {
		index = 0
	}
	if index >= len(sortedReturns) {
		index = len(sortedReturns) - 1
	}

	varReturn := sortedReturns[index]
	// VaR is typically expressed as a positive number representing loss
	return -varReturn * portfolioValue
}
