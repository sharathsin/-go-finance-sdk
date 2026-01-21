package money

import (
	"fmt"
	"strings"

	"github.com/shopspring/decimal"
)

// Money represents a monetary amount in a specific currency.
// It uses decimal.Decimal for high precision arithmetic.
type Money struct {
	amount   decimal.Decimal
	currency string
}

// New creates a new Money instance.
func New(amount decimal.Decimal, currency string) Money {
	return Money{
		amount:   amount,
		currency: strings.ToUpper(currency),
	}
}

// NewFromFloat creates a new Money instance from a float64.
func NewFromFloat(amount float64, currency string) Money {
	return New(decimal.NewFromFloat(amount), currency)
}

// NewFromString creates a new Money instance from a string.
func NewFromString(amount string, currency string) (Money, error) {
	dec, err := decimal.NewFromString(amount)
	if err != nil {
		return Money{}, err
	}
	return New(dec, currency), nil
}

// Amount returns the decimal amount.
func (m Money) Amount() decimal.Decimal {
	return m.amount
}

// Currency returns the currency code.
func (m Money) Currency() string {
	return m.currency
}

// Add adds two Money instances. Returns error if currencies mismatch.
func (m Money) Add(other Money) (Money, error) {
	if m.currency != other.currency {
		return Money{}, fmt.Errorf("currency mismatch: %s vs %s", m.currency, other.currency)
	}
	return New(m.amount.Add(other.amount), m.currency), nil
}

// Sub subtracts other Money from m. Returns error if currencies mismatch.
func (m Money) Sub(other Money) (Money, error) {
	if m.currency != other.currency {
		return Money{}, fmt.Errorf("currency mismatch: %s vs %s", m.currency, other.currency)
	}
	return New(m.amount.Sub(other.amount), m.currency), nil
}

// Mul multiplies Money by a scalar.
func (m Money) Mul(scalar decimal.Decimal) Money {
	return New(m.amount.Mul(scalar), m.currency)
}

// Div divides Money by a scalar.
func (m Money) Div(scalar decimal.Decimal) Money {
	return New(m.amount.Div(scalar), m.currency)
}

// String returns the string representation.
func (m Money) String() string {
	return fmt.Sprintf("%s %s", m.amount.StringFixed(2), m.currency)
}

// IsZero returns true if amount is zero.
func (m Money) IsZero() bool {
	return m.amount.IsZero()
}

// IsPositive returns true if amount is positive.
func (m Money) IsPositive() bool {
	return m.amount.IsPositive()
}

// IsNegative returns true if amount is negative.
func (m Money) IsNegative() bool {
	return m.amount.IsNegative()
}
