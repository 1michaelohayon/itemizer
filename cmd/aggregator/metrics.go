package main

import (
	"1michaelohayon/itemizer/typ"
	"fmt"

	"github.com/prometheus/client_golang/prometheus"
)

type Metrics struct {
	errCounterItem  prometheus.Counter
	errCounterKafka prometheus.Counter
	kafkaMessages   prometheus.Counter
	StorageUnits    map[Identifier]*prometheus.Counter
	Sender          map[Identifier]*prometheus.Counter
}

func NewGlobalMetrics() *Metrics {
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
		StorageUnits: make(map[Identifier]*prometheus.Counter),
		Sender:       make(map[Identifier]*prometheus.Counter),
	}
	prometheus.MustRegister(m.kafkaMessages)
	prometheus.MustRegister(m.errCounterItem)
	prometheus.MustRegister(m.errCounterKafka)
	return m
}

type Identifier string

func MetricStorageUnitId(data typ.ItemData) Identifier {
	return Identifier(fmt.Sprintf("%s", data.StorageUnit.ID))
}

func MetricSenderId(data typ.ItemData) Identifier {
	return Identifier(fmt.Sprintf("%s_%d", data.StorageUnit.ID, data.Item.Sender.ID))
}
