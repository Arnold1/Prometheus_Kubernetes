package main

import (
	"net/http"
	"time"
	"math/rand"
	"github.com/prometheus/client_golang/prometheus"
	"fmt"
)

// YourHandler displays the string below
func YourHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("this is not the endpoint you are looking for... try /metrics"))
}

func doSomeWork() int {
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)

	statusCodes := [...]int{
		http.StatusOK,
		http.StatusBadRequest,
		http.StatusUnauthorized,
		http.StatusInternalServerError,
	}
	return statusCodes[rand.Intn(len(statusCodes))]
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	var status int

	defer func() {
		// hello_requests_total{status="200"} 2385
		counter.With(prometheus.Labels{
			"status": fmt.Sprint(status),
		}).Inc()
    }()

    status = doSomeWork()
    w.WriteHeader(status)
    w.Write([]byte("Hello, World!\n"))
}