package instrument

// InstrumentType defines the type of financial instrument.
type InstrumentType string

const (
    TypeEquity InstrumentType = "EQUITY"
    TypeBond   InstrumentType = "BOND"
    TypeOption InstrumentType = "OPTION"
)

// Instrument represents a tradeable financial asset.
type Instrument interface {
    ID() string
    Type() InstrumentType
    Currency() string
}
