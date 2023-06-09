package main

import (
	"fmt"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
)

type HTTPfunc func(http.ResponseWriter, *http.Request) error

type HTTPMetricHandler struct {
	errCounter prometheus.Counter
}

func newHTTPMetricsHandler(reqName string) *HTTPMetricHandler {
	errCounter := prometheus.NewCounter(prometheus.CounterOpts{
		Namespace: fmt.Sprintf("http_%s_%s", reqName, "err_counter"),
		Name:      "item",
	})

	return &HTTPMetricHandler{
		errCounter: errCounter,
	}
}
