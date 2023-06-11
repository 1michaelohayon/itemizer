package main

import (
	"1michaelohayon/itemizer/typ"
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/sirupsen/logrus"
)

var client = &http.Client{}

func informAPi(itemD typ.ItemData) {
	json, err := json.Marshal(itemD)
	if err != nil {
		logrus.Error("marshal error:", err)
	}
	req, err := http.NewRequest("POST", ApiEndPoint, bytes.NewBuffer(json))
	if err != nil {
		logrus.Error("new request error:", err)
	}
	req.Header.Set("Content-Type", "application/json")
	_, err = client.Do(req)
	if err != nil {
		logrus.Error("clent req send error:", err)
	}
}
