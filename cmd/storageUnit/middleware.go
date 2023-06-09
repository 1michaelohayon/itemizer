package main

import (
	"1michaelohayon/itemizer/typ"
	"time"

	"github.com/sirupsen/logrus"
)

type LogMiddleware struct {
	next DataProducer
}

func NewLogMiddleware(next DataProducer) *LogMiddleware {
	return &LogMiddleware{
		next: next,
	}
}

func (l *LogMiddleware) ProduceData(data typ.ItemData) error {
	defer func(start time.Time) {
		logrus.WithFields(logrus.Fields{
			"Item ID":        data.Item.ID,
			"StorageUnit ID": data.StorageUnit.ID,
			"Sender ID":      data.Item.Sender.ID,
			"took":           time.Since(start),
		}).Info("Producing to Kafka")
	}(time.Now())
	return l.next.ProduceData(data)
}
