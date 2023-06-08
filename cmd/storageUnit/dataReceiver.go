package main

import (
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
		p           DataProducer
		err         error
		kafakaTopic = "Itemizer"
	)
	p, err = NewKafkaProducer(kafakaTopic)
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
		var data typ.Item
		if err := dr.wsCon.ReadJSON(&data); err != nil {
			log.Println("read error:", err)
			readErrors++
			if readErrors > 5 {
				dr.wsCon.Close()
				break
			}
			continue
		}
		fmt.Printf("<-- Received item:%d from sender:%d\n", data.ID, data.Sender.ID)

		if err := dr.kProd.ProduceData(data); err != nil {
			fmt.Println("kafka ProduceData error:", err)
		}
	}
}
