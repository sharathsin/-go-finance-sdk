package pricing

import (
	"math"
	"math/rand"
	"sync"
	"time"

	"github.com/antigravity/go-finance-sdk/pkg/instrument"
)

// MonteCarloPricer implements Monte Carlo simulation for pricing.
type MonteCarloPricer struct {
	Simulations  int
	RiskFreeRate float64
	Volatility   float64
}

// NewMonteCarloPricer creates a new Monte Carlo pricer.
func NewMonteCarloPricer(sims int, r, sigma float64) *MonteCarloPricer {
	return &MonteCarloPricer{
		Simulations:  sims,
		RiskFreeRate: r,
		Volatility:   sigma,
	}
}

// Price calculates the price using concurrent simulations.
func (mc *MonteCarloPricer) Price(inst instrument.Instrument) (float64, error) {
	opt, ok := inst.(*instrument.Option)
	if !ok {
		return 0, nil
	}

	S0 := 100.0 // Placeholder spot
	K, _ := opt.Strike().Float64()
	T := time.Until(opt.Expiry()).Hours() / (24 * 365)
	r := mc.RiskFreeRate
	sigma := mc.Volatility

	// Parallel processing
	numGoroutines := 10
	simsPerRoutine := mc.Simulations / numGoroutines
	results := make(chan float64, numGoroutines)
	var wg sync.WaitGroup

	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			rng := rand.New(rand.NewSource(time.Now().UnixNano()))
			sumPayoff := 0.0
			for j := 0; j < simsPerRoutine; j++ {
				z := rng.NormFloat64()
				ST := S0 * math.Exp((r-0.5*sigma*sigma)*T+sigma*math.Sqrt(T)*z)
				payoff := 0.0
				if opt.OptionType() == instrument.Call {
					if ST > K {
						payoff = ST - K
					}
				} else {
					if K > ST {
						payoff = K - ST
					}
				}
				sumPayoff += payoff
			}
			results <- sumPayoff
		}()
	}

	wg.Wait()
	close(results)

	totalPayoff := 0.0
	for p := range results {
		totalPayoff += p
	}

	averagePayoff := totalPayoff / float64(mc.Simulations)
	discountedPrice := averagePayoff * math.Exp(-r*T)

	return discountedPrice, nil
}
