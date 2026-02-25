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
	//из докера поднят контейнер с кафкой:
	//21c4cab0a70e:/$ /opt/kafka/bin/kafka-topics.sh --create \
	//  --bootstrap-server localhost:9092 \
	//  --replication-factor 1 \
	//  --partitions 1 \
	//  --topic test-topic
	//Created topic test-topic.
	//21c4cab0a70e:/$ /opt/kafka/bin/kafka-topics.sh --describe \
	//  --topic test-topic
	//or: 1   Configs: min.insync.replicas=1,segment.bytes=1073741824
	//        Topic: test-topic       Partition: 0    Leader: 1       Replicas: 1     Isr: 1  Elr:  LastKnownElr:

	config := sarama.NewConfig()
	config.Version = sarama.V2_1_0_0
	config.Producer.Return.Successes = true

	producer, err := sarama.NewAsyncProducer([]string{broker}, config)
	if err != nil {
		log.Fatalf("error with producer: %v", err)
	}
	defer producer.AsyncClose()

	consumer, err := sarama.NewConsumer([]string{broker}, config)
	if err != nil {
		log.Fatalf("error with consumer: %v", err)
	}
	//defer consumer.Close()

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
		//_, _, err := producer.SendMessage(&sarama.ProducerMessage{
		msg := &sarama.ProducerMessage{
			Topic: topic,
			Value: sarama.StringEncoder(value),
		}
		select {

		case producer.Input() <- msg:
			log.Printf("sent async message: %s", value)
		case err := <-producer.Errors():
			log.Fatalf("async producer error: %v", err)
			//if err != nil {
			//	log.Fatalf("error with sending message: %v", err)
			//}
			//log.Printf("sent message: %s", value)
		}

	}
	<-done
	log.Println("done!")
}
