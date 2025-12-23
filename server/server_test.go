package server

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthCheck(t *testing.T) {
	tests := []struct {
		name           string
		method         string
		expectedStatus int
		expectedBody   HealthResponse
	}{
		{
			name:           "GET /health",
			method:         http.MethodGet,
			expectedStatus: http.StatusOK,
			expectedBody: HealthResponse{
				Status:  "ok",
				Message: "server works!",
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(tt.method, "/health", nil)

			rr := httptest.NewRecorder()

			HealthCheck(rr, req)

			if status := rr.Code; status != tt.expectedStatus {
				t.Errorf("handler вернул статус %d, ожидался %d", status, tt.expectedStatus)
			}

			var got HealthResponse
			if err := json.NewDecoder(rr.Body).Decode(&got); err != nil {
				t.Fatalf("не удалось распарсить JSON: %v", err)
			}

			if got != tt.expectedBody {
				t.Errorf("handler вернул %+v, ожидался %+v", got, tt.expectedBody)
			}
		})
	}
}