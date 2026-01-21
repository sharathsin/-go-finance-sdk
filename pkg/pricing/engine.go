package pricing

import (
    "github.com/antigravity/go-finance-sdk/pkg/instrument"
)

// Pricer interface for pricing instruments.
type Pricer interface {
    Price(inst instrument.Instrument) (float64, error)
}
