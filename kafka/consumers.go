package kafka

import (
	"context"
	"encoding/json"
	"log"

	"github.com/segmentio/kafka-go"
)

func readerInit(topic_name string, groupBalancers []kafka.GroupBalancer) *kafka.Reader {
	config := kafka.ReaderConfig{
		Brokers:        []string{KAFKA_ADDR},
		Topic:          topic_name,
		GroupID:        "payment-consumer",
		MinBytes:       10e3, // 10KB
		MaxBytes:       10e6, // 10MB
		GroupBalancers: groupBalancers,
	}
	r := kafka.NewReader(config)
	return r
}

func readerRead(r *kafka.Reader) {
	defer r.Close()
	for {
		msg, err := r.ReadMessage(context.Background())
		if err != nil {
			log.Fatal("could not read message: " + err.Error())
			break
		}
		var val MessageInterface
		err = json.Unmarshal(msg.Value, &val)

		if err != nil {
			log.Fatalf("Failed to unmarshal message: %v", err.Error())
			continue
		}
		// fmt.Printf("message at topic [%v] partition [%v] offset [%v]: %s - %s\n", msg.Topic, msg.Partition, msg.Offset, string(msg.Key), string(msg.Value))
	}
}

func messageRead(topic_name string) {
	groupBalancers := make([]kafka.GroupBalancer, 0)
	groupBalancers = append(groupBalancers, kafka.RangeGroupBalancer{})

	readers := make([]*kafka.Reader, 0)
	for i := 0; i < PARTITION; i++ {
		readers = append(readers, readerInit(topic_name, groupBalancers))
	}
	for _, reader := range readers {
		go readerRead(reader)
	}
}
