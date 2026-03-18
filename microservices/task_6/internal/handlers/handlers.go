package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"task_6/internal/models"
)

var (
	mu    sync.RWMutex
	users = make(map[string]models.User)
)

func UsersHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case http.MethodPost:
		registerUser(w, r)
	case http.MethodGet:
		listUsers(w, r)
	case http.MethodPut, http.MethodDelete:
		updateOrDeleteUser(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func registerUser(w http.ResponseWriter, r *http.Request) {
	mu.RLock()
	defer mu.RUnlock()

	var u models.User
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	id := strconv.Itoa(len(users) + 1)
	u.ID = id
	users[id] = u

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(u) //nolint:errcheck
}

func listUsers(w http.ResponseWriter, r *http.Request) {
	mu.RLock()
	defer mu.RUnlock()

	list := make([]models.User, 0, len(users))
	for _, u := range users {
		list = append(list, u)
	}
	json.NewEncoder(w).Encode(list) //nolint:errcheck
}

func updateOrDeleteUser(w http.ResponseWriter, r *http.Request) {
	mu.RLock()
	defer mu.RUnlock()

	path := strings.Trim(r.URL.Path, "/")
	parts := strings.Split(path, "/")

	if len(parts) != 2 || parts[0] != "users" {
		http.Error(w, "wrong path", http.StatusBadRequest)
		return
	}
	id := parts[1]

	if _, exists := users[id]; !exists && r.Method == http.MethodPut {
		http.Error(w, "user is not found", http.StatusNotFound)
		return
	}

	if r.Method == http.MethodPut {
		var u models.User
		if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
			http.Error(w, "JSON error!", http.StatusBadRequest)
			return
		}
		u.ID = id
		users[id] = u
		json.NewEncoder(w).Encode(u) //nolint:errcheck
		return
	}

	delete(users, id)
	w.WriteHeader(http.StatusNoContent)
}
