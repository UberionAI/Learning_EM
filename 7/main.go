package main

import (
	"log"

	"github.com/IBM/sarama"
)

const (
	broker = "localhost:9092"
	topic  = "test-topic"
)

func main() {
	config := sarama.NewConfig()
	config.Version = sarama.V2_1_0_0
	config.Producer.Return.Successes = true

	producer, err := sarama.NewSyncProducer([]string{broker}, config)
	if err != nil {
		log.Fatalf("producer: %v", err)
	}
	defer func() { _ = producer.Close() }()

	messages := []struct {
		key   string
		value string
	}{
		{"user-1", "event-1 for user-1"},
		{"user-1", "event-2 for user-1"},
		{"user-2", "event-1 for user-2"},
		{"user-3", "event-1 for user-3"},
	}

	for _, m := range messages {
		msg := &sarama.ProducerMessage{
			Topic: topic,
			Key:   sarama.StringEncoder(m.key),
			Value: sarama.StringEncoder(m.value),
		}

		partition, offset, err := producer.SendMessage(msg)
		if err != nil {
			log.Fatalf("send with key=%s: %v", m.key, err)
		}

		log.Printf("Sent key=%q value=%q to partition=%d offset=%d",
			m.key, m.value, partition, offset)
	}
}
