package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {

	uptimeGauge := prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: "example_app",
		Name:      "uptime_seconds",
		Help:      "How long the system has been up for",
	})
	prometheus.MustRegister(uptimeGauge)

	timer := time.NewTimer(time.Second)
	go func() {
		for {
			<-timer.C
			uptimeGauge.Inc()
			timer = time.NewTimer(time.Second)
		}
	}()

	http.HandleFunc(
		"/hello",
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "Hello!  20200304!!!!!!!!!!!")
		},
	)

	http.Handle("/metrics", promhttp.Handler())

	log.Fatal(http.ListenAndServe(":8080", nil))
}
