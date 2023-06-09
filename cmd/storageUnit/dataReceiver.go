package main

import (
	"1michaelohayon/itemizer/config"
	"1michaelohayon/itemizer/typ"
	"fmt"
	"log"

	"github.com/gorilla/websocket"
)

type DataReceiver struct {
	msgch chan typ.Item
	wsCon *websocket.Conn
	kProd DataProducer
}

func NewDataReceiver() (*DataReceiver, error) {
	var (
		p   DataProducer
		err error
	)

	p, err = NewKafkaProducer(config.KafkaTopic, config.KafkaHost)
	if err != nil {
		return nil, err
	}
	p = NewLogMiddleware(p)

	return &DataReceiver{
		msgch: make(chan typ.Item, 128),
		kProd: p,
	}, nil
}

func (dr *DataReceiver) wsReceiveLoop() {
	readErrors := 0
	fmt.Println("New Scanner connected.")
	for {
		var item typ.Item
		if err := dr.wsCon.ReadJSON(&item); err != nil {
			log.Println("read error:", err)
			readErrors++
			if readErrors > 5 {
				dr.wsCon.Close()
				break
			}
			continue
		}
		fmt.Printf("<-- Received item:%d from sender:%d\n", item.ID, item.Sender.ID)

		data := typ.ItemData{
			StorageUnit: this,
			Item:        item,
		}
		if err := dr.kProd.ProduceData(data); err != nil {
			fmt.Println("kafka ProduceData error:", err)
		}
	}
}
