package kafka

import (
	"bytes"
	"context"
	"encoding/json"
	"time"

	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/snappy"
)

func PushMessage(value *CreateOrderMessageInterface, topic string, partition int) (bool, error) {
	config := kafka.WriterConfig{
		Brokers:          []string{KAFKA_ADDR},
		Topic:            topic,
		Balancer:         &PartitionBalancer{},
		WriteTimeout:     10 * time.Second,
		ReadTimeout:      10 * time.Second,
		CompressionCodec: snappy.NewCompressionCodec(),
	}
	writer_connector := kafka.NewWriter(config)
	defer writer_connector.Close()

	new_byte_buffer := new(bytes.Buffer)
	json.NewEncoder(new_byte_buffer).Encode(value)

	err := writer_connector.WriteMessages(
		context.Background(),
		kafka.Message{
			Partition: partition,
			Value:     new_byte_buffer.Bytes(),
		},
	)
	if err != nil {
		return false, err
	}
	return true, nil
}
