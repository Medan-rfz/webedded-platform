package middleware

import (
	"fmt"
	"net/http"
	"regexp"
	"time"

	"webedded.users_management/internal/app/metrics"
)

type metricsMiddleware struct {
	handler http.Handler
}

func MetricsMiddleware(next http.Handler) http.Handler {
	return &metricsMiddleware{next}
}

func (l *metricsMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	recorder := &statusRecorder{
		ResponseWriter: w,
		Status:         200,
	}

	l.handler.ServeHTTP(recorder, r)
	duration := time.Since(start)

	status := fmt.Sprintf("%v", recorder.Status)
	re := regexp.MustCompile(`\d+`)
	cleanUrl := re.ReplaceAllString(r.URL.Path, "{}")

	metrics.MetricRequestsCounterInc(status, cleanUrl, r.Method)
	metrics.MetricResponceTime(duration, status, cleanUrl, r.Method)
}
