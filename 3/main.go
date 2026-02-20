package main

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/IBM/sarama"
)

const (
	broker     = "localhost:9092"
	topic      = "test-topic"
	numWorkers = 5
)

func main() {
	config := sarama.NewConfig()
	config.Version = sarama.V2_1_0_0

	producer, err := sarama.NewAsyncProducer([]string{broker}, config)
	if err != nil {
		log.Fatalf("producer: %v", err)
	}
	defer func() {
		producer.AsyncClose()
		for range producer.Errors() {
		}
		for range producer.Successes() {
		}
	}()

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

	messages := make(chan *sarama.ConsumerMessage, 100)
	var wg sync.WaitGroup

	for w := 0; w < numWorkers; w++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			for msg := range messages {
				time.Sleep(100 * time.Millisecond)
				log.Printf("worker-%d: offset=%d value=%s",
					workerID, msg.Offset, string(msg.Value))
			}
		}(w)
	}

	done := make(chan struct{})
	go func() {
		defer close(messages)
		count := 0
		for msg := range pc.Messages() {
			messages <- msg
			count++
			if count == 10 {
				return
			}
		}
	}()

	for i := 0; i < 10; i++ {
		value := fmt.Sprintf("message-%d", i)
		msg := &sarama.ProducerMessage{Topic: topic, Value: sarama.StringEncoder(value)}
		select {
		case producer.Input() <- msg:
			log.Printf("produced: %s", value)
		case perr := <-producer.Errors():
			log.Fatalf("producer error: %v", perr)
		}
	}

	<-done
	wg.Wait()
	log.Println("Multi-threaded consumer done!")
}
