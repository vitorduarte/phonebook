package prometheus

import (
	"net/http"
	"strconv"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/vitorduarte/phonebook/internal/interceptor"
)

var requestsByMethodAndPath = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "http_requests",
		Help: "Mumber of HTTP requests made by each method in paths.",
	},
	[]string{"method", "path"},
)

var responseStatus = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "http_responses_status",
		Help: "Number of HTTP responses by status code.",
	},
	[]string{"status"},
)

var httpDurations = prometheus.NewHistogramVec(
	prometheus.HistogramOpts{
		Name: "http_response_time_seconds",
		Help: "Duration of HTTP requests.",
	},
	[]string{"method", "path"},
)

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		wr := interceptor.NewResponseRecorder(w)
		timer := prometheus.NewTimer(httpDurations.WithLabelValues(r.Method, r.URL.Path))
		next.ServeHTTP(wr, r)

		requestsByMethodAndPath.WithLabelValues(r.Method, r.URL.Path).Inc()
		responseStatus.WithLabelValues(strconv.Itoa(wr.Status)).Inc()
		timer.ObserveDuration()
	})
}

func Init() {
	prometheus.Register(requestsByMethodAndPath)
	prometheus.Register(responseStatus)
	prometheus.Register(httpDurations)
}
