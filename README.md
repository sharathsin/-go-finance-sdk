# Go Finance SDK

A comprehensive Go library for financial engineering, offering tools for money management, instrument modeling, pricing engines, and risk management.

## Features

- **Money & Currency**: High-precision arithmetic using `decimal` type, currency support.
- **Instruments**: Support for Equities, Bonds, and Options (European/American).
- **Pricing Engines**:
  - Black-Scholes Model
  - Monte Carlo Simulation
- **Risk Management**: Historical Value at Risk (VaR) calculation.

## Installation

```bash
go get github.com/antigravity/go-finance-sdk
```

## Usage

### Money Arithmetic

```go
package main

import (
	"fmt"
	"github.com/antigravity/go-finance-sdk/pkg/money"
)

func main() {
	m1 := money.NewFromFloat(100.50, "USD")
	m2 := money.NewFromFloat(50.25, "USD")
	
	sum, _ := m1.Add(m2)
	fmt.Println(sum) // 150.75 USD
}
```

### Option Pricing (Black-Scholes)

```go
package main

import (
	"fmt"
	"time"
	"github.com/shopspring/decimal"
	"github.com/antigravity/go-finance-sdk/pkg/instrument"
	"github.com/antigravity/go-finance-sdk/pkg/pricing"
)

func main() {
	// Setup
	underlying := instrument.NewEquity("AAPL", "USD", "AAPL")
	strike := decimal.NewFromInt(150)
	expiry := time.Now().Add(30 * 24 * time.Hour)
	
	opt := instrument.NewEuropeanOption("OPT1", underlying, strike, expiry, instrument.Call)
	
	// Pricing
	// Rate=5%, Volatility=20%
	pricer := pricing.NewBlackScholesPricer(0.05, 0.20)
	price, _ := pricer.Price(opt)
	
	fmt.Printf("Option Price: %.2f\n", price)
}
```

### Value at Risk (VaR)

```go
package main

import (
	"fmt"
	"github.com/antigravity/go-finance-sdk/pkg/risk"
)

func main() {
	returns := []float64{-0.02, 0.01, -0.01, 0.03, -0.05}
	portfolioValue := 10000.0
	confidence := 0.95
	
	varValue := risk.CalculateHistoricalVaR(returns, confidence, portfolioValue)
	fmt.Printf("VaR (95%%): %.2f\n", varValue)
}
```

## Testing

Run the test suite:

```bash
go test ./...
```
