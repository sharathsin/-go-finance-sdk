package main

import (
	"fmt"
	"time"

	"github.com/shopspring/decimal"
	"github.com/antigravity/go-finance-sdk/pkg/instrument"
	"github.com/antigravity/go-finance-sdk/pkg/pricing"
	"github.com/antigravity/go-finance-sdk/pkg/risk"
)

func main() {
	fmt.Println("Starting Financial SDK Simulation...")

	// 1. Create an Option
	strike := decimal.NewFromFloat(100.0)
	expiry := time.Now().AddDate(0, 3, 0) // 3 months
	underlying := instrument.NewEquity("AAPL", "USD", "AAPL")
	
	callOption := instrument.NewEuropeanOption("OPT-1", underlying, strike, expiry, instrument.Call)

	fmt.Printf("Instrument: %s (%s) Strike: %s\n", callOption.ID(), callOption.OptionType(), callOption.Strike())

	// 2. Price using Black-Scholes
	bsPricer := pricing.NewBlackScholesPricer(0.05, 0.2) // r=5%, sigma=20%
	bsPrice, _ := bsPricer.Price(callOption)
	fmt.Printf("Black-Scholes Price: %.4f\n", bsPrice)

	// 3. Price using Monte Carlo (Concurrent)
	mcPricer := pricing.NewMonteCarloPricer(100000, 0.05, 0.2)
	mcPrice, _ := mcPricer.Price(callOption)
	fmt.Printf("Monte Carlo Price: %.4f (100k simulations)\n", mcPrice)

	// 4. Calculate Risk (VaR)
	// Mock historical returns
	returns := []float64{-0.02, -0.01, 0.005, 0.01, 0.015, -0.03, 0.02}
	portfolioValue := 100000.0
	var95 := risk.CalculateHistoricalVaR(returns, 0.95, portfolioValue)
	fmt.Printf("Historical VaR (95%%): %.2f\n", var95)
}
