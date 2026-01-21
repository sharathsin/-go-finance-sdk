package pricing

import (
    "math"
    "time"

    "github.com/antigravity/go-finance-sdk/pkg/instrument"
)

// BlackScholesPricer implements the Black-Scholes pricing model.
type BlackScholesPricer struct {
    RiskFreeRate float64
    Volatility   float64
}

// NewBlackScholesPricer creates a new Black-Scholes pricer.
func NewBlackScholesPricer(r, sigma float64) *BlackScholesPricer {
    return &BlackScholesPricer{
        RiskFreeRate: r,
        Volatility:   sigma,
    }
}

func (bs *BlackScholesPricer) Price(inst instrument.Instrument) (float64, error) {
    opt, ok := inst.(*instrument.Option)
    if !ok {
        // Simple stub: only options supported for this example logic
        return 0, nil
    }

    S := 100.0 // Placeholder spot price since we don't have market data feed yet
    K, _ := opt.Strike().Float64()
    T := time.Until(opt.Expiry()).Hours() / (24 * 365)
    r := bs.RiskFreeRate
    sigma := bs.Volatility

    d1 := (math.Log(S/K) + (r+sigma*sigma/2.0)*T) / (sigma * math.Sqrt(T))
    d2 := d1 - sigma*math.Sqrt(T)

    if opt.OptionType() == instrument.Call {
        return S*normCdf(d1) - K*math.Exp(-r*T)*normCdf(d2), nil
    }
    // Put
    return K*math.Exp(-r*T)*normCdf(-d2) - S*normCdf(-d1), nil
}

// Standard normal cumulative distribution function
func normCdf(x float64) float64 {
    return 0.5 * (1 + math.Erf(x/math.Sqrt2))
}
