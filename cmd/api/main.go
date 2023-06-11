package main

import (
	"1michaelohayon/itemizer/cmd/api/db"
	"fmt"
	"log"
	"net/http"
	"os"
)

var (
	listenPort = ":5000"
)

func init() {
	envPort := os.Getenv("API_LSTEN_PORT")

	if len(envPort) > 0 {
		listenPort = envPort
	}
	db.ConnectToPSQL()
	db.CreateItemDataTable()
}

func main() {
	fmt.Println("Listening on port", listenPort)

	http.HandleFunc("/", ItemsApi)
	log.Fatal(http.ListenAndServe(listenPort, nil))
}
