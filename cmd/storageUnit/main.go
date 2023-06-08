package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

func main() {
	fmt.Println("Awaiting Connections. . .")

	go http.HandleFunc("/ws", handleWS) //start connection on a new gorotuine
	http.ListenAndServe(":30000", nil)
}

func handleWS(w http.ResponseWriter, r *http.Request) {
	u := websocket.Upgrader{
		ReadBufferSize:  1028,
		WriteBufferSize: 1028,
	}
	conn, err := u.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}

	rec, err := NewDataReceiver()
	if err != nil {
		log.Fatal(err)
	}
	rec.wsCon = conn
	rec.wsReceiveLoop()
}
