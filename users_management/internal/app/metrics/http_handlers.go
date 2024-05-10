package metrics

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var requestsCounter = promauto.NewCounterVec(
	prometheus.CounterOpts{
		Name: "http_requests_total",
		Help: "Total number of HTTP requests",
	}, []string{"code", "handler", "method"})

var responseTime = promauto.NewHistogramVec(
	prometheus.HistogramOpts{
		Name:    "http_request_duration_seconds",
		Help:    "Response time duration and status code",
		Buckets: prometheus.DefBuckets,
	}, []string{"code", "handler", "method"})

func MetricRequestsCounterInc(code, handler, method string) {
	requestsCounter.WithLabelValues(code, handler, method).Inc()
}

func MetricResponceTime(duration time.Duration, code, handler, method string) {
	responseTime.WithLabelValues(code, handler, method).Observe(float64(duration.Seconds()))
}
