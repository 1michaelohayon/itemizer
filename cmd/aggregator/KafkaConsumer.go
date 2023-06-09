package main

import (
	"1michaelohayon/itemizer/config"
	"1michaelohayon/itemizer/typ"
	"encoding/json"
	"fmt"
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
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
		Metrics:  NewMetrics(),
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
		logrus.Errorf("JSON serialization error: %s\n")
	}
	fmt.Println("Consumed", data)
	errs := len(data.Item.Errors)
	if errs > 0 {
		c.Metrics.errCounterItem.Add(float64(errs))
	}

	c.Metrics.kafkaMessages.Inc()
	return nil
}
