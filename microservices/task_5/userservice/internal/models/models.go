package models

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type RegistrationEvent struct {
	UserID string `json:"user_id"`
	Name   string `json:"name"`
}
