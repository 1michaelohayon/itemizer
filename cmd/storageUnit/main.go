package main

import (
	"1michaelohayon/itemizer/cmd/storageUnit/db"
	"1michaelohayon/itemizer/typ"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/websocket"
)

var this = typ.StorageUnit{
	ID: "placeholder",
}

func init() {
	id := os.Getenv("STORAGE_UNIT_ID")
	if len(id) > 0 {
		this.ID = id
	} else {
		//panic/logfatal in production
		fmt.Println("ID was not given")
	}

	if err := db.Connect("./data.db"); err != nil {
		log.Fatal(err)
	}
	db.CreateItemsTable()
}

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
