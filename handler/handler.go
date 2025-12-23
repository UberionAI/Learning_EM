package main

import (
	"encoding/json"
	"net/http"
)

type HealthResponse struct {
	Alive bool `json:"alive"`
}

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(HealthResponse{Alive: true})
}
