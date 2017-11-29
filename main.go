// this project is heavily inspired by the prometheus golang client exmaple project
// https://github.com/prometheus/client_golang/blob/master/examples/simple/main.go

package main

import (
	"net/http"
	"os"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	promhttp "github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
)

var (
	// Counter vector to which we can attach labels. That creates many key-value
    // label combinations. So in our case we count requests by status code separetly.
    counter = prometheus.NewCounterVec(
        prometheus.CounterOpts{
            Name: "hello_requests_total",
            Help: "Total number of /hello requests.",
        },
        []string{"status"},
    )
)

func init() {
    prometheus.MustRegister(counter)
}

// ExampleGauge which will increment the incValue gauge by 1 every 5 seconds and promote it to /metrics
func ExampleGauge() {
	incValue := prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: "acme_company",
		Subsystem: "blob_storage",
		Name:      "inc_value",
		Help:      "just an increased number.",
	})
	// register incValue in promtthp handler
	prometheus.Register(incValue)
	// loop over the ticker and call Inc function
	for range time.Tick(time.Second * 5) {

		// increment incValue by 1 every 5 seconds
		incValue.Inc()
	}
}

func main() {
	var log = logrus.New()
	// set the log output
	log.Out = os.Stdout
	// set the log level
	log.Level = logrus.DebugLevel
	// Routes consist of a path and a handler function.
	r := mux.NewRouter()
	// sample log.Info
	log.Info("http server is ready")
	// sample log.Debug
	log.Debug("i am only visible in debug mode\n")
	// exposes / endpoint with the YourHandler handler
	r.HandleFunc("/", YourHandler)
	// exposes /hello endpoint with the helloHandler handler
	r.HandleFunc("/hello", helloHandler)
	// exposes /metrics endpoint with standard golang metrics used by prometheus
	r.Handle("/metrics", promhttp.Handler())
	// wrap a logger arount the mux router
	loggedRouter := handlers.LoggingHandler(os.Stdout, r)
	// wrap a prometheus instrument handler around the logger
	prometheusRouter := prometheus.InstrumentHandlerWithOpts(prometheus.SummaryOpts{}, loggedRouter)

	// start a goroutine which start the polling for the metrics endpoint
 	go ExampleGauge()

	// Bind to a port and pass our loggedRouter in
	log.Fatal(http.ListenAndServe(":8080", prometheusRouter))

}
