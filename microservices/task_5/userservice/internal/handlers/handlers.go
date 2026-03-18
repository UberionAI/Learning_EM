package handlers

import (
	"encoding/json"
	"github.com/segmentio/kafka-go"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"userservice/internal/models"
)

var (
	users = make(map[string]models.User)
	mu    sync.Mutex
)

const kafkaBroker = "127.0.0.1:9092"
const topic = "user-registrations"

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
	mu.Lock()
	defer mu.Unlock()

	var u models.User
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	id := strconv.Itoa(len(users) + 1)
	u.ID = id
	users[id] = u

	writer := &kafka.Writer{
		Addr:     kafka.TCP(kafkaBroker),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	}
	defer writer.Close()
	msg := models.RegistrationEvent{UserID: id, Name: u.Name}
	if err := writer.WriteMessages(r.Context(), kafka.Message{Value: []byte(toJSON(msg))}); err != nil {
		http.Error(w, "Kafka error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(u)
}

func listUsers(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	list := make([]models.User, 0, len(users))
	for _, u := range users {
		list = append(list, u)
	}
	json.NewEncoder(w).Encode(list)
}

func updateOrDeleteUser(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	path := strings.Trim(r.URL.Path, "/")
	parts := strings.Split(path, "/")
	if len(parts) != 2 || parts[0] != "users" {
		http.Error(w, "wrong path", http.StatusBadRequest)
		return
	}
	id := parts[1]
	if _, exists := users[id]; !exists {
		if r.Method == http.MethodPut {
			http.Error(w, "user is not found", http.StatusNotFound)
			return
		}
		delete(users, id)
		w.WriteHeader(http.StatusNoContent)
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
		json.NewEncoder(w).Encode(u)
	}
}

func toJSON(v interface{}) string {
	b, _ := json.Marshal(v)
	return string(b)
}
