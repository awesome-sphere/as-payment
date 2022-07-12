package kafka

import (
	"github.com/segmentio/kafka-go"
)

var kafkaLeader *kafka.Conn

func InitializeKafka() {
	topic := "payment"

	conn, err := kafka.Dial("tcp", "localhost:9095")
	if err != nil {
		panic(err.Error())
	}
	defer conn.Close()

	topicConfigs := []kafka.TopicConfig{
		{
			Topic:             topic,
			NumPartitions:     5,
			ReplicationFactor: 1,
		},
	}

	err = kafkaLeader.CreateTopics(topicConfigs...)
	if err != nil {
		panic(err.Error())
	}

	kafkaLeader = conn
}
