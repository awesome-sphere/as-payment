package kafka

import (
	"github.com/segmentio/kafka-go"
)

func ListTopic(connector *kafka.Conn) map[string]*TopicInterface {
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

func isTopicExist(connector *kafka.Conn, topic_name string) bool {
	topic_map := ListTopic(connector)
	_, ok := topic_map[topic_name]
	return ok
}

func InitializeKafka() {
	topic := "payment"

	conn, err := kafka.Dial("tcp", "localhost:9092")
	if err != nil {
		panic(err.Error())
	}
	defer conn.Close()

	if !isTopicExist(conn, topic) {
		topicConfigs := []kafka.TopicConfig{
			{
				Topic:             topic,
				NumPartitions:     5,
				ReplicationFactor: 1,
			},
		}

		err := conn.CreateTopics(topicConfigs...)
		if err != nil {
			panic(err.Error())
		}
	}
}
