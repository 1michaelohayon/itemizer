package main

import (
	"1michaelohayon/itemizer/typ"
	"fmt"
	"log"
	"math"
	"math/rand"
	"os"
	"time"

	"github.com/gorilla/websocket"
)

var (
	wsEndPoint = "ws://127.0.0.1:30000/ws"
	senderId   = rand.Intn(math.MaxInt - 1)
)

func init() {
	ep := os.Getenv("SCANNER_ENDPOINT")
	if len(ep) > 0 {
		wsEndPoint = ep
	}
	fmt.Println("Connecting to -->", wsEndPoint)
}

func main() {
	fmt.Println("Connecting to", wsEndPoint)
	conn, _, err := websocket.DefaultDialer.Dial(wsEndPoint, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected.")
	for {
		item := NewRndItem()
		if err := conn.WriteJSON(item); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Sent %+v -->\n", item)
		time.Sleep(5 * time.Second)
	}
}

func NewRndItem() typ.Item {
	id := rand.Intn(8) + 1
	return typ.Item{
		ID:     int64(id),
		Name:   fmt.Sprintf("Random Item %d", id),
		Amount: 1,
		Sender: typ.Sender{
			ID:   int64(senderId),
			Time: time.Now(),
		},
	}
}
