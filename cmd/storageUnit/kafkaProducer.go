package main

import (
	"1michaelohayon/itemizer/typ"
	"encoding/json"
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type DataProducer interface {
	ProduceData(typ.ItemData) error
}

type KafkaProducer struct {
	producer *kafka.Producer
	topic    string
}

func NewKafkaProducer(topic, host string) (DataProducer, error) {
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": host})
	if err != nil {
		return nil, err
	}

	// stat another goroutine to check if succesfully delivered the data
	go func() {
		for e := range p.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Printf("Delivery failed: %v\n", ev.TopicPartition)
				} else {
					fmt.Printf("Delivered message to %v\n", ev.TopicPartition)
				}
			}
		}
	}()

	return &KafkaProducer{
		producer: p,
		topic:    topic,
	}, nil
}

func (p *KafkaProducer) ProduceData(data typ.ItemData) error {
	b, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return p.producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &p.topic,
			Partition: kafka.PartitionAny,
		},
		Value: b,
	}, nil)
}
