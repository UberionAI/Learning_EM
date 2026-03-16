package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"Learning_EM/task_1/internal/handlers"
)

func BenchmarkListUsersREST(b *testing.B) {
	req, _ := http.NewRequest("GET", "/users", nil)
	rr := httptest.NewRecorder()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		handlers.UsersHandler(rr, req)
		rr.Body.Reset()
	}
}
