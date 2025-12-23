package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthCheckHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/health", nil)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(HealthCheckHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var got, want HealthResponse
	json.NewDecoder(rr.Body).Decode(&got)
	want = HealthResponse{Alive: true}

	if got != want {
		t.Errorf("handler returned unexpected body: got %v want %v", got, want)
	}
}

//PS C:\Users\tla\GolandProjects\Learning_EM\handler> go test -v .\...
//=== RUN   TestHealthCheckHandler
//--- PASS: TestHealthCheckHandler (0.00s)
//PASS
//ok      github.com/UberionAI/Learning_EM/handler        (cached)
//PS C:\Users\tla\GolandProjects\Learning_EM\handler>
