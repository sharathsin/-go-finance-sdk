package money

import (
	"testing"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	m := New(decimal.NewFromInt(100), "USD")
	assert.Equal(t, "100", m.Amount().String())
	assert.Equal(t, "USD", m.Currency())
}

func TestNewFromFloat(t *testing.T) {
	m := NewFromFloat(100.50, "USD")
	assert.Equal(t, "100.5", m.Amount().String())
	assert.Equal(t, "USD", m.Currency())
}

func TestNewFromString(t *testing.T) {
	m, err := NewFromString("100.50", "USD")
	assert.NoError(t, err)
	assert.Equal(t, "100.5", m.Amount().String())
	assert.Equal(t, "USD", m.Currency())

	_, err = NewFromString("invalid", "USD")
	assert.Error(t, err)
}

func TestAdd(t *testing.T) {
	m1 := NewFromFloat(100, "USD")
	m2 := NewFromFloat(50, "USD")

	sum, err := m1.Add(m2)
	assert.NoError(t, err)
	assert.Equal(t, "150", sum.Amount().String())

	m3 := NewFromFloat(50, "EUR")
	_, err = m1.Add(m3)
	assert.Error(t, err)
}

func TestSub(t *testing.T) {
	m1 := NewFromFloat(100, "USD")
	m2 := NewFromFloat(50, "USD")

	diff, err := m1.Sub(m2)
	assert.NoError(t, err)
	assert.Equal(t, "50", diff.Amount().String())

	m3 := NewFromFloat(50, "EUR")
	_, err = m1.Sub(m3)
	assert.Error(t, err)
}

func TestMul(t *testing.T) {
	m := NewFromFloat(100, "USD")
	res := m.Mul(decimal.NewFromFloat(1.5))
	assert.Equal(t, "150", res.Amount().String())
}

func TestDiv(t *testing.T) {
	m := NewFromFloat(100, "USD")
	res := m.Div(decimal.NewFromInt(2))
	assert.Equal(t, "50", res.Amount().String())
}

func TestString(t *testing.T) {
	m := NewFromFloat(100.5, "USD")
	assert.Equal(t, "100.50 USD", m.String())
}
