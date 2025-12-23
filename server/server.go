package server

import (
	"encoding/json"
	"net/http"
)

type HealthResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	resp := HealthResponse{
		Status:  "ok",
		Message: "server works!",
	}
	json.NewEncoder(w).Encode(resp)
}

//PS C:\Users\uberion\GolandProjects\Learning_EM\server> go test
//PASS
//ok      github.com/UberionAI/Learning_EM/server 0.421s
//PS C:\Users\uberion\GolandProjects\Learning_EM\server>