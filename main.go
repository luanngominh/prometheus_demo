package main

import (
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	xxCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "page_counter",
			Help: "Number of request access to xx page",
		},
		[]string{"page"},
	)
)

func init() {
	prometheus.MustRegister(xxCounter)
}

func main() {
	http.Handle("/metrics", promhttp.Handler())

	http.HandleFunc("/xxx", func(w http.ResponseWriter, q *http.Request) {
		xxCounter.With(prometheus.Labels{"page": "xx_page"}).Inc()
		w.Write([]byte("meow meow meow"))
	})
	http.HandleFunc("/yyy", func(w http.ResponseWriter, q *http.Request) {
		xxCounter.With(prometheus.Labels{"page": "yy_page"}).Inc()
		w.Write([]byte("meow meow meow"))
	})
	log.Fatal(http.ListenAndServe(":9876", nil))
}
