package observability

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	// HttpRequestDuration measures the duration of HTTP requests.
	HttpRequestDuration = promauto.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "finance_sdk_http_request_duration_seconds",
		Help:    "Duration of HTTP requests to external APIs",
		Buckets: prometheus.DefBuckets,
	}, []string{"provider", "method", "status"})

	// ExternalApiErrors counts the number of errors from external APIs.
	ExternalApiErrors = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "finance_sdk_external_api_errors_total",
		Help: "Total number of errors encountered when calling external APIs",
	}, []string{"provider", "type"})

	// CircuitBreakerState tracks the state of circuit breakers.
	CircuitBreakerState = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "finance_sdk_circuit_breaker_state",
		Help: "Current state of the circuit breaker (0: Closed, 1: Open, 2: Half-Open)",
	}, []string{"name"})
)
