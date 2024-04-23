package middleware

import (
	"log"
	"net/http"
	"time"
)

type statusRecorder struct {
	http.ResponseWriter
	Status int
}

type logger struct {
	handler http.Handler
}

func LoggerMiddleware(next http.Handler) http.Handler {
	return &logger{next}
}

func (l *logger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	recorder := &statusRecorder{
		ResponseWriter: w,
		Status:         200,
	}

	start := time.Now()
	l.handler.ServeHTTP(recorder, r)
	duration := time.Since(start).Milliseconds()

	log.Printf("Request details:\n"+
		"  Protocol: http\n"+
		"  Method: %s\n"+
		"  Path: %s\n"+
		"  Duration: %v ms\n"+
		"  Status Code: %v %s\n",
		r.Method, r.RequestURI, duration, recorder.Status, http.StatusText(recorder.Status))
}

func (r *statusRecorder) WriteHeader(status int) {
	r.Status = status
	r.ResponseWriter.WriteHeader(status)
}
