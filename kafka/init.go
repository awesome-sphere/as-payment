package kafka

import (
	"bytes"
	"context"
	"encoding/json"
	"strconv"
	"time"

	"github.com/awesome-sphere/as-payment/utils"
	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/snappy"
)

var TOPIC string
var PARTITION int

func PushMessage(topic string, n_partition string, value *MessageInterface) {
	config := kafka.WriterConfig{
		Brokers:          []string{"localhost:9092"},
		Topic:            topic,
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
			Key:   []byte(n_partition),
			Value: new_byte_buffer.Bytes(),
		},
	)
	if err != nil {
		panic(err.Error())
	}
}

func ListTopics(connector *kafka.Conn) map[string]*TopicInterface {
	partitions, err := connector.ReadPartitions()
	if err != nil {
		panic(err.Error())
	}
	m := make(map[string]*TopicInterface)
	for _, p := range partitions {
		if _, ok := m[p.Topic]; ok {
			m[p.Topic].Partition += 1
		} else {
			m[p.Topic] = &TopicInterface{Partition: 1}
		}
	}
	return m
}

func isExisted(connector *kafka.Conn, topic string) bool {
	topic_map := ListTopics(connector)
	_, ok := topic_map[topic]
	return ok
}

func InitializeKafka() {
	TOPIC := utils.GetenvOr("KAFKA_TOPIC", "payment")
	PARTITION, err := strconv.Atoi(utils.GetenvOr("KAFKA_TOPIC_PARTITION", "5"))
	if err != nil {
		panic(err.Error())
	}

	conn, err := kafka.Dial("tcp", "localhost:9092")
	if err != nil {
		panic(err.Error())
	}
	defer conn.Close()

	if !isExisted(conn, TOPIC) {
		topicConfigs := []kafka.TopicConfig{
			{
				Topic:             TOPIC,
				NumPartitions:     PARTITION,
				ReplicationFactor: 1,
			},
		}

		err := conn.CreateTopics(topicConfigs...)
		if err != nil {
			panic(err.Error())
		}
	}
}