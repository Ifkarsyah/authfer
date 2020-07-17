package main

import (
	"encoding/json"
	"net/http"
)

func ResponseOK(w http.ResponseWriter, body interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	return json.NewEncoder(w).Encode(body)
}
