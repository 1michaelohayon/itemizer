package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	httpPort = ":4000"
)

func init() {
	hp := os.Getenv("AGG_PORT")
	if len(hp) > 0 {
		httpPort = hp
	}
}

func main() {
	kafkaConsumer, err := NewKafkaConsumer()
	if err != nil {
		log.Fatal(err)
	}

	go kafkaConsumer.Start() // new go routine

	log.Fatal(NewHttpListener(httpPort))
}

func NewHttpListener(listenAddr string) error {
	fmt.Println("Http port running on", listenAddr)
	http.Handle("/metrics", promhttp.Handler())
	return http.ListenAndServe(listenAddr, nil)

}

type Aggregator interface {
	Consumer() (*KafkaConsumer, error)
}
