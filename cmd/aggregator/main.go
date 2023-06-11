package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	httpPort    = ":4000"
	ApiEndPoint = "http://localhost:5000"
)

func init() {
	hp := os.Getenv("AGG_PORT")
	if len(hp) > 0 {
		httpPort = hp
	}
	ep := os.Getenv("HTTP_ENDPOINT")
	if len(ep) > 0 {
		ApiEndPoint = ep
	}
}

func main() {
	time.Sleep(15 * time.Second)
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
