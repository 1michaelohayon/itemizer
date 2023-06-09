package main

import (
	"github.com/prometheus/client_golang/prometheus"
)

type Metrics struct {
	errCounterItem  prometheus.Counter
	errCounterKafka prometheus.Counter
	kafkaMessages   prometheus.Counter
}

func NewMetrics() *Metrics {
	m := &Metrics{
		errCounterItem: prometheus.NewCounter(
			prometheus.CounterOpts{
				Namespace: "item_error_counter",
				Name:      "item",
			}),
		errCounterKafka: prometheus.NewCounter(
			prometheus.CounterOpts{
				Namespace: "kafka_error_counter",
				Name:      "kafka",
			}),
		kafkaMessages: prometheus.NewCounter(
			prometheus.CounterOpts{
				Namespace: "kafka_messages_counter",
				Name:      "kafka",
			}),
	}
	prometheus.MustRegister(m.kafkaMessages)
	prometheus.MustRegister(m.errCounterItem)
	prometheus.MustRegister(m.errCounterKafka)
	return m
}
