package metrics

import "github.com/prometheus/client_golang/prometheus"

func init() {
	prometheus.MustRegister(HttpRequestsDuration)
	prometheus.MustRegister(TotalWins)
	prometheus.MustRegister(TotalLosses)
}

var HttpRequestsDuration = prometheus.NewHistogramVec(prometheus.HistogramOpts{
	Name: "http_request_duration_microseconds",
	Help: "Duration of all HTTP requests",
}, []string{"status", "method", "path"})

var TotalWins = prometheus.NewCounterVec(prometheus.CounterOpts{
	Name: "total_wins",
	Help: "The total number of games won by snake",
}, []string{"snakeID"})

var TotalLosses = prometheus.NewCounterVec(prometheus.CounterOpts{
	Name: "total_loses",
	Help: "The total number of games losed by snake",
}, []string{"snakeID"})
