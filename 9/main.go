package main

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/IBM/sarama"
)

const (
	broker  = "localhost:9092"
	topic   = "test-topic"
	groupID = "ack-group"
)

type keyValueHandler struct{}

func (h *keyValueHandler) Setup(s sarama.ConsumerGroupSession) error {
	log.Println("consumer group session setup, memberID:", s.MemberID())
	return nil
}

func (h *keyValueHandler) Cleanup(s sarama.ConsumerGroupSession) error {
	log.Println("consumer group session cleanup")
	return nil
}

func (h *keyValueHandler) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		log.Printf("received: key=%q value=%q partition=%d offset=%d",
			string(msg.Key), string(msg.Value), msg.Partition, msg.Offset)

		sess.MarkMessage(msg, "")

		time.Sleep(100 * time.Millisecond)
	}
	return nil
}

func main() {
	cfg := sarama.NewConfig()
	cfg.Version = sarama.V2_1_0_0
	cfg.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRange
	cfg.Consumer.Offsets.Initial = sarama.OffsetOldest

	group, err := sarama.NewConsumerGroup([]string{broker}, groupID, cfg)
	if err != nil {
		log.Fatalf("new consumer group: %v", err)
	}
	defer func() { _ = group.Close() }()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	handler := &keyValueHandler{}
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

	time.Sleep(10 * time.Second)
	cancel()
	wg.Wait()
	log.Println("consumer with acks stopped")
}
