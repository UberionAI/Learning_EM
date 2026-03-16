package main

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
	"task_6/internal/handlers"
	"time"
)

var (
	requestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: "userservice",
			Subsystem: "http",
			Name:      "request_duration_seconds",
			Help:      "HTTP request duration by handler",
			Buckets:   prometheus.DefBuckets,
		},
		[]string{"handler", "method"},
	)
	responseSize = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: "userservice",
			Subsystem: "http",
			Name:      "response_size_bytes",
			Help:      "HTTP response size",
			Buckets:   prometheus.ExponentialBuckets(100, 2, 10),
		},
		[]string{"handler"},
	)
)

func init() {
	prometheus.MustRegister(requestDuration)
	prometheus.MustRegister(responseSize)
}

type sizeTrackingResponseWriter struct {
	http.ResponseWriter
	size int
}

func (rw *sizeTrackingResponseWriter) Write(b []byte) (int, error) {
	size, err := rw.ResponseWriter.Write(b)
	rw.size += size
	return size, err
}

func instrumentedHandler(handlerName string, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		rw := &sizeTrackingResponseWriter{ResponseWriter: w}
		next.ServeHTTP(rw, r)
		requestDuration.WithLabelValues(handlerName, r.Method).Observe(time.Since(start).Seconds())
		responseSize.WithLabelValues(handlerName).Observe(float64(rw.size))
	}
}

func main() {
	http.HandleFunc("/users", instrumentedHandler("users", handlers.UsersHandler))
	http.HandleFunc("/users/", instrumentedHandler("users", handlers.UsersHandler))
	http.Handle("/metrics", promhttp.Handler())
	fmt.Println("Listening on port 8080")
	http.ListenAndServe(":8080", nil) //nolint:errcheck
}
