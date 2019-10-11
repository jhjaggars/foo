package main

import (
	"fmt"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var counter = promauto.NewCounter(prometheus.CounterOpts{
	Name: "foo_requests",
	Help: "counter of http requests",
})

func main() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("got called")
		counter.Inc()
		w.WriteHeader(http.StatusOK)
	}
	http.HandleFunc("/", handler)
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":8080", nil)
}
