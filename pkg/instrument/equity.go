package instrument

// Equity represents a stock or share.
type Equity struct {
    id       string
    currency string
    symbol   string
}

// NewEquity creates a new Equity instrument.
func NewEquity(id, currency, symbol string) *Equity {
    return &Equity{
        id:       id,
        currency: currency,
        symbol:   symbol,
    }
}

func (e *Equity) ID() string {
    return e.id
}

func (e *Equity) Type() InstrumentType {
    return TypeEquity
}

func (e *Equity) Currency() string {
    return e.currency
}

func (e *Equity) Symbol() string {
    return e.symbol
}
