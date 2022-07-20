package kafka

import (
	"bytes"
	"context"
	"encoding/json"
	"time"

	"github.com/awesome-sphere/as-payment/kafka/interfaces"
	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/snappy"
)

func pushMessage(topic string, partition int) *kafka.Writer {
	config := kafka.WriterConfig{
		Brokers:          []string{KAFKA_ADDR},
		Topic:            topic,
		Balancer:         &PartitionBalancer{},
		WriteTimeout:     10 * time.Second,
		ReadTimeout:      10 * time.Second,
		CompressionCodec: snappy.NewCompressionCodec(),
	}
	writer_connector := kafka.NewWriter(config)

	return writer_connector
}

func UpdateTopic(value *interfaces.UpdateOrderMessageInterface, topic string, partition int) (bool, error) {
	writer_connector := pushMessage(topic, partition)
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

func CreateTopic(value *interfaces.CreateOrderMessageInterface, topic string, partition int) (bool, error) {
	writer_connector := pushMessage(topic, partition)
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
