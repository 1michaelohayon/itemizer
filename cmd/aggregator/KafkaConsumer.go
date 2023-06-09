package main

import (
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
}

func NewKafkaConsumer(topic, host string) (*KafkaConsumer, error) {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": host, //docker network
		"group.id":          "myGroup",
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		panic(err)
	}

	if err = c.SubscribeTopics([]string{topic}, nil); err != nil {
		return nil, err
	}

	return &KafkaConsumer{
		consumer: c,
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
		msg, err := c.consumer.ReadMessage(-1)
		if err != nil {
			//TODO: inc in promehtheus
			logrus.Errorf("kafka consumer error: %s\n", err)
			log.Fatal(err)
		}
		var data typ.ItemData
		if err := json.Unmarshal(msg.Value, &data); err != nil {
			//TODO: inc in promehtheus
			logrus.Errorf("JSON serialization error: %s\n")
			continue
		}

		fmt.Println("Consumed", data)
		//.... TODO

	}
}
