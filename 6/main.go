package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/IBM/sarama"
)

const (
	broker = "localhost:9092"
	topic  = "test-topic"
)

func main() {
	config := sarama.NewConfig()
	config.Version = sarama.V2_1_0_0

	consumer, err := sarama.NewConsumer([]string{broker}, config)
	if err != nil {
		log.Fatalf("consumer: %v", err)
	}
	defer func() { _ = consumer.Close() }()

	pc, err := consumer.ConsumePartition(topic, 0, sarama.OffsetOldest)
	if err != nil {
		log.Fatalf("partition consumer: %v", err)
	}
	defer func() { _ = pc.Close() }()

	log.Printf("Console consumer started. Topic=%s partition=0", topic)
	log.Println("Press Ctrl+C to exit.")

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, syscall.SIGTERM)

	for {
		select {
		case msg := <-pc.Messages():
			log.Printf("message: partition=%d offset=%d value=%s",
				msg.Partition, msg.Offset, string(msg.Value))
		case <-signals:
			log.Println("Stopping console consumer...")
			return
		}
	}
}
