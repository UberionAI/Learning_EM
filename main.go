package main

import (
	"fmt"
	"github.com/IBM/sarama"
	"log"
	"time"
)

const (
	broker = "localhost:9092"
	topic  = "test-topic"
)

func main() {
	config := sarama.NewConfig()
	config.Version = sarama.V2_1_0_0
	config.Producer.Return.Successes = true

	admin, err := sarama.NewClusterAdmin([]string{broker}, config)
	if err != nil {
		log.Fatalf("error with admin: %v", err)
	}
	defer admin.Close()

	err = admin.CreateTopic(topic, &sarama.TopicDetail{
		NumPartitions:     1,
		ReplicationFactor: 1,
	}, false)
	if err != nil {
		log.Printf("topic already exists or error: %v", err)
	}

	producer, err := sarama.NewSyncProducer([]string{broker}, config)
	if err != nil {
		log.Fatalf("error with producer: %v", err)
	}
	defer producer.Close()

	consumer, err := sarama.NewConsumer([]string{broker}, config)
	if err != nil {
		log.Fatalf("error with consumer: %v", err)
	}
	defer consumer.Close()

	//start partition 0
	partitions, err := consumer.ConsumePartition(topic, 0, sarama.OffsetOldest)
	if err != nil {
		log.Fatalf("error with consumer: %v", err)
	}
	defer partitions.Close()

	//starting reading in goroutine
	done := make(chan struct{})
	go func() {
		defer close(done)
		count := 0
		for msg := range partitions.Messages() {
			log.Printf("consumed message: offset=%d value=%s", msg.Offset, string(msg.Value))
			count++
			if count == 10 {
				return
			}
		}
	}()

	time.Sleep(500 * time.Millisecond)

	for i := 0; i < 10; i++ {
		value := fmt.Sprintf("message-%d", i)
		_, _, err := producer.SendMessage(&sarama.ProducerMessage{
			Topic: topic,
			Value: sarama.StringEncoder(value),
		})
		if err != nil {
			log.Fatalf("error with sending message: %v", err)
		}
		log.Printf("sent message: %s", value)
	}
	<-done
	log.Println("done!")
}
