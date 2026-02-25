package main

import (
	"context"
	"fmt"
	"github.com/IBM/sarama"
	"log"
	"sync"
	"time"
)

const (
	broker = "localhost:9092"
	topic  = "test-topic"
)

func main() {
	config := sarama.NewConfig()
	config.Version = sarama.V2_1_0_0
	config.Consumer.Offsets.Initial = sarama.OffsetOldest

	producer, err := sarama.NewAsyncProducer([]string{broker}, config)
	if err != nil {
		log.Fatalf("producer: %v", err)
	}
	defer producer.AsyncClose()

	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())

	for partition := 0; partition < 3; partition++ {
		wg.Add(1)
		go func(p int) {
			defer wg.Done()
			consumer, err := sarama.NewConsumer([]string{broker}, config)
			if err != nil {
				log.Printf("consumer %d: %v", p, err)
				return
			}
			defer consumer.Close()

			pc, err := consumer.ConsumePartition(topic, int32(p), sarama.OffsetOldest)
			if err != nil {
				log.Printf("partition %d: %v", p, err)
				return
			}
			defer pc.Close()

			log.Printf("Goroutine partition=%d started", p)
			for {
				select {
				case msg := <-pc.Messages():
					log.Printf("P%d consumed: offset=%d value=%s", p, msg.Offset, string(msg.Value))
				case <-ctx.Done():
					return
				}
			}
		}(partition)
	}

	time.Sleep(2 * time.Second)
	for i := 0; i < 10; i++ {
		value := fmt.Sprintf("message-%d", i)
		producer.Input() <- &sarama.ProducerMessage{
			Topic: topic,
			Value: sarama.StringEncoder(value),
		}
		log.Printf("produced: %s", value)
	}

	time.Sleep(5 * time.Second)
	cancel()
	wg.Wait()
	log.Println(" multi-consumer done!")
}
