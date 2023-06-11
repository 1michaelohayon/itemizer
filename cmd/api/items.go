package main

import (
	"1michaelohayon/itemizer/cmd/api/db"
	"1michaelohayon/itemizer/typ"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func ItemsApi(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case http.MethodGet:
		GetAll(w, r)
	case http.MethodPost:
		AddItemData(w, r)
	}

}

func AddItemData(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest)
	}
	defer r.Body.Close()

	var data *typ.ItemData
	err = json.Unmarshal(body, &data)
	if err != nil {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}

	if err := db.InsertItem(*data); err != nil {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	} else {
		http.Error(w, "ok", http.StatusNoContent)
	}

}

func GetAll(w http.ResponseWriter, r *http.Request) {
	itemsData, err := db.SelectAll()
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
	jsonData, err := json.Marshal(itemsData)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
	w.Write(jsonData)
}
