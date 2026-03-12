package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"task_2/internal/models"
)

var Users = make(map[string]models.User)

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

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
	var u models.User
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	id := strconv.Itoa(len(Users) + 1)
	u.ID = id
	Users[id] = u

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(u)
}

func listUsers(w http.ResponseWriter, r *http.Request) {
	list := make([]models.User, 0, len(Users))
	for _, u := range Users {
		list = append(list, u)
	}
	json.NewEncoder(w).Encode(list)
}

func updateOrDeleteUser(w http.ResponseWriter, r *http.Request) {
	path := strings.Trim(r.URL.Path, "/")
	parts := strings.Split(path, "/")

	if len(parts) != 2 || parts[0] != "users" {
		http.Error(w, "wrong path", http.StatusBadRequest)
		return
	}
	id := parts[1]

	if _, exists := Users[id]; !exists && r.Method == http.MethodPut {
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
		Users[id] = u
		json.NewEncoder(w).Encode(u)
		return
	}

	delete(Users, id)
	w.WriteHeader(http.StatusNoContent)
}
