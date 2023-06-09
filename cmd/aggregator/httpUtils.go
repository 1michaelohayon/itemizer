package main

import (
	"encoding/json"
	"net/http"
)

func writeJSON(w http.ResponseWriter, status int, v any) error {
	w.WriteHeader(status)
	w.Header().Add("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(v)
}

type ApiError struct {
	Code int
	Err  error
}

// tostring
func (e ApiError) Error() string {
	return e.Err.Error()
}

func NewHTTPHandler(fn HTTPfunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := fn(w, r); err != nil {
			if apiErr, ok := err.(ApiError); ok {
				writeJSON(w, apiErr.Code, map[string]string{"error": apiErr.Error()})
			}
		}
	}
}
