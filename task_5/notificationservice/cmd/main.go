package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/segmentio/kafka-go"
	"log"
	"notificationservice/internal/models"
)

const broker = "127.0.0.1:9092"
const topic = "user-registrations"
const groupID = "notification-group"

func main() {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{broker},
		Topic:   topic,
		GroupID: groupID,
	})
	defer reader.Close()

	fmt.Println("NotificationService started, consuming from", topic)
	for {
		msg, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Printf("Read error: %v", err)
			continue
		}
		var event models.RegistrationEvent
		if err := json.Unmarshal(msg.Value, &event); err != nil {
			log.Printf("JSON error: %v", err)
			continue
		}

		fmt.Printf("Уведомление пользователю %s (ID: %s): Добро пожаловать!\n", event.Name, event.UserID)
	}
}
