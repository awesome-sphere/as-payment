package kafka

import (
	"bytes"
	"context"
	"encoding/json"
	"strconv"
	"time"

	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/snappy"
)

func PushMessage(value *MessageInterface) (bool, error) {
	config := kafka.WriterConfig{
		Brokers:          []string{KAFKA_ADDR},
		Topic:            TOPIC,
		Balancer:         &kafka.LeastBytes{},
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
			Key:   []byte(strconv.Itoa(value.TheaterId)),
			Value: new_byte_buffer.Bytes(),
		},
	)
	if err != nil {
		return false, err
	}
	return true, nil
}

func SubmitToConsumer(user_id int, time_slot_id int, theater_id int, seat_number []int, price int, duration time.Time) {
	PushMessage(&MessageInterface{
		UserID:     user_id,
		TimeSlotId: time_slot_id,
		TheaterId:  theater_id,
		SeatNumber: seat_number,
		Price:      price,
		Duration:   time.Now(),
	})
}
