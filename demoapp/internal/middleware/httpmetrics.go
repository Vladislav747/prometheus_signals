package middleware

import (
	"github.com/go-chi/chi/v5"
	"net/http"
	"prometheus_signals/internal/helpers"
	"prometheus_signals/internal/metrics"
	"time"
)

func HTTPMetrics(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		srw := helpers.NewStatusResponseWriter(w)
		now := time.Now()

		next.ServeHTTP(srw, r)

		elapsedSeconds := time.Since(now).Seconds()
		pattern := chi.RouteContext(r.Context()).RoutePattern()
		method := chi.RouteContext(r.Context()).RouteMethod
		status := srw.GetStatusString()

		metrics.HttpRequestsTotal.WithLabelValues(pattern, method, status).Inc()
		metrics.HttpRequestsDurationHistorgram.WithLabelValues(pattern, method).Observe(elapsedSeconds)
		metrics.HttpRequestsDurationSummary.WithLabelValues(pattern, method).Observe(elapsedSeconds)
	})
}
