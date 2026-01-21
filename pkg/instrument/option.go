package instrument

import (
    "time"

    "github.com/shopspring/decimal"
)

type OptionType string

const (
    Call OptionType = "CALL"
    Put  OptionType = "PUT"
)

type ExerciseStyle string

const (
    European ExerciseStyle = "EUROPEAN"
    American ExerciseStyle = "AMERICAN"
)

// Option represents a financial option contract.
type Option struct {
    id            string
    underlying    Instrument
    strike        decimal.Decimal
    expiry        time.Time
    optionType    OptionType
    exerciseStyle ExerciseStyle
}

// NewEuropeanOption creates a new European option.
func NewEuropeanOption(id string, underlying Instrument, strike decimal.Decimal, expiry time.Time, optType OptionType) *Option {
    return &Option{
        id:            id,
        underlying:    underlying,
        strike:        strike,
        expiry:        expiry,
        optionType:    optType,
        exerciseStyle: European,
    }
}

func (o *Option) ID() string {
    return o.id
}

func (o *Option) Type() InstrumentType {
    return TypeOption
}

func (o *Option) Currency() string {
    return o.underlying.Currency()
}

func (o *Option) Underlying() Instrument {
    return o.underlying
}

func (o *Option) Strike() decimal.Decimal {
    return o.strike
}

func (o *Option) Expiry() time.Time {
    return o.expiry
}

func (o *Option) OptionType() OptionType {
    return o.optionType
}

func (o *Option) Style() ExerciseStyle {
    return o.exerciseStyle
}
