package main

import (
	"1michaelohayon/itemizer/config"
	"1michaelohayon/itemizer/typ"
	"encoding/json"
	"fmt"
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/sirupsen/logrus"
)

type KafkaConsumer struct {
	consumer  *kafka.Consumer
	isRunning bool
	Metrics   *Metrics
}

func NewKafkaConsumer() (*KafkaConsumer, error) {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": config.KafkaHost, //docker network
		"group.id":          "myGroup",
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		panic(err)
	}

	if err = c.SubscribeTopics([]string{config.KafkaTopic}, nil); err != nil {
		return nil, err
	}

	return &KafkaConsumer{
		consumer: c,
		Metrics:  NewGlobalMetrics(),
	}, nil
}

func (c *KafkaConsumer) Close() {
	c.isRunning = false
}

func (c *KafkaConsumer) Start() {
	logrus.Info("Kafka transport started")
	c.isRunning = true
	c.readMessageLoop()
}

func (c *KafkaConsumer) readMessageLoop() {
	for c.isRunning {
		c.ConsumeData()
	}
}

func (c *KafkaConsumer) ConsumeData() error {
	msg, err := c.consumer.ReadMessage(-1)
	if err != nil {
		c.Metrics.errCounterKafka.Inc()
		logrus.Errorf("kafka consumer error: %s\n", err)
		log.Fatal(err)
	}
	var data typ.ItemData
	if err := json.Unmarshal(msg.Value, &data); err != nil {
		c.Metrics.errCounterKafka.Inc()
		logrus.Errorf("JSON serialization error: %s\n", err)
	}
	fmt.Println("Consumed", data)
	errs := len(data.Item.Errors)
	if errs > 0 {
		c.Metrics.errCounterItem.Add(float64(errs))
	}

	c.StorageUnitMetricHandle(data)
	c.SenderMetricHandle(data)

	c.Metrics.kafkaMessages.Inc()
	return nil
}

func (c *KafkaConsumer) StorageUnitMetricHandle(data typ.ItemData) {
	ptr := c.Metrics.StorageUnits[MetricStorageUnitId(data)]
	if ptr != nil {
		(*ptr).Inc()
	} else {
		counter := prometheus.NewCounter(
			prometheus.CounterOpts{
				Namespace: "storage_unitss",
				Name:      data.StorageUnit.ID,
			})

		prometheus.MustRegister(counter)
		c.Metrics.StorageUnits[MetricStorageUnitId(data)] = &counter
	}
}

func (c *KafkaConsumer) SenderMetricHandle(data typ.ItemData) {
	ptr := c.Metrics.Sender[MetricSenderId(data)]
	if ptr != nil {
		(*ptr).Inc()
	} else {
		counter := prometheus.NewCounter(
			prometheus.CounterOpts{
				Namespace: fmt.Sprintf("sender_from_%s", data.StorageUnit.ID),
				Name:      fmt.Sprintf("senderId:%d", data.Item.Sender.ID),
			})
		prometheus.MustRegister(counter)
		c.Metrics.Sender[MetricSenderId(data)] = &counter
	}
}
