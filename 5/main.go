package main

import (
	"fmt"
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
	config.Producer.Partitioner = sarama.NewManualPartitioner

	producer, err := sarama.NewSyncProducer([]string{broker}, config)
	if err != nil {
		log.Fatalf("producer: %v", err)
	}
	defer func() { _ = producer.Close() }()

	for partition := int32(0); partition < 3; partition++ {
		msg := &sarama.ProducerMessage{
			Topic:     topic,
			Partition: partition,
			Value:     sarama.StringEncoder(fmt.Sprintf("msg-to-partition-%d", partition)),
		}

		p, offset, err := producer.SendMessage(msg)
		if err != nil {
			log.Fatalf("send to partition %d: %v", partition, err)
		}
		log.Printf("Sent to partition=%d offset=%d", p, offset)
	}
}
