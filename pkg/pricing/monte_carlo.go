package pricing

import (
	"context"
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
	rngPool      sync.Pool
}

// NewMonteCarloPricer creates a new Monte Carlo pricer.
func NewMonteCarloPricer(sims int, r, sigma float64) *MonteCarloPricer {
	return &MonteCarloPricer{
		Simulations:  sims,
		RiskFreeRate: r,
		Volatility:   sigma,
		rngPool: sync.Pool{
			New: func() interface{} {
				return rand.New(rand.NewSource(time.Now().UnixNano()))
			},
		},
	}
}

// Price calculates the price using concurrent simulations.
func (mc *MonteCarloPricer) Price(ctx context.Context, inst instrument.Instrument) (float64, error) {
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

			// Check context before starting
			if ctx.Err() != nil {
				return
			}

			// Get RNG from pool
			rng := mc.rngPool.Get().(*rand.Rand)
			defer mc.rngPool.Put(rng)

			sumPayoff := 0.0
			for j := 0; j < simsPerRoutine; j++ {
				// Check context periodically (every 1000 sims or so) to avoid overhead
				if j%1000 == 0 && ctx.Err() != nil {
					return
				}

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

	// fast exit if cancelled
	if ctx.Err() != nil {
		return 0, ctx.Err()
	}

	totalPayoff := 0.0
	for p := range results {
		totalPayoff += p
	}

	averagePayoff := totalPayoff / float64(mc.Simulations)
	discountedPrice := averagePayoff * math.Exp(-r*T)

	return discountedPrice, nil
}
