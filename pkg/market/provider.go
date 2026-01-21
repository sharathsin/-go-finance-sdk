package market

import (
	"context"
	"time"

	"github.com/shopspring/decimal"
)

// Price represents the price of an instrument at a specific time.
type Price struct {
	Symbol    string
	Value     decimal.Decimal
	Timestamp time.Time
}

// Provider defines the interface for fetching market data.
type Provider interface {
	GetPrice(ctx context.Context, symbol string) (Price, error)
}
