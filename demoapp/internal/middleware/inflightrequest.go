package middleware

import (
	"net/http"
	"prometheus_signals/internal/metrics"
)

func InflightRequests(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		metrics.HttpRequestsCurrent.WithLabelValues().Inc()

		next.ServeHTTP(w, r)

		metrics.HttpRequestsCurrent.WithLabelValues().Dec()
	})
}
