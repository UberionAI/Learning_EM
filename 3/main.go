package main

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/IBM/sarama"
)

const (
	broker  = "localhost:9092"
	topic   = "test-topic"
	groupID = "test-group"
)

type messageHandler struct {
	processed int
	mu        sync.Mutex
}

func (h *messageHandler) Setup(_ sarama.ConsumerGroupSession) error   { return nil }
func (h *messageHandler) Cleanup(_ sarama.ConsumerGroupSession) error { return nil }
func (h *messageHandler) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		h.mu.Lock()
		h.processed++
		h.mu.Unlock()

		log.Printf("group consumed: offset=%d value=%s", msg.Offset, string(msg.Value))
		sess.MarkMessage(msg, "")
	}
	return nil
}

func main() {
	config := sarama.NewConfig()
	config.Version = sarama.V2_1_0_0
	config.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRoundRobin
	config.Consumer.Offsets.Initial = sarama.OffsetOldest

	producer, err := sarama.NewAsyncProducer([]string{broker}, config)
	if err != nil {
		log.Fatalf("producer: %v", err)
	}
	defer producer.AsyncClose()

	group, err := sarama.NewConsumerGroup([]string{broker}, groupID, config)
	if err != nil {
		log.Fatalf("consumer group: %v", err)
	}
	defer func() { _ = group.Close() }()

	handler := &messageHandler{}
	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		for {
			if err := group.Consume(ctx, []string{topic}, handler); err != nil {
				log.Printf("consume error: %v", err)
			}
			if ctx.Err() != nil {
				return
			}
		}
	}()

	time.Sleep(2 * time.Second)

	// Producer
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

	time.Sleep(5 * time.Second)
	cancel()
	wg.Wait()
	log.Printf("Group consumer processed %d messages!", handler.processed)
}
