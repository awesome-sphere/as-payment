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
var KAFKA_ADDR string

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
	TOPIC = utils.GetenvOr("KAFKA_TOPIC", "payment")
	var err error
	PARTITION, err = strconv.Atoi(utils.GetenvOr("KAFKA_TOPIC_PARTITION", "5"))
	KAFKA_ADDR = utils.GetenvOr("KAFKA_ADDR", "localhost:9092")
	if err != nil {
		panic(err.Error())
	}

	conn, err := kafka.Dial("tcp", KAFKA_ADDR)
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
