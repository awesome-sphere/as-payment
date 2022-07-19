package kafka

import (
	"strconv"

	"github.com/awesome-sphere/as-payment/utils"
	"github.com/segmentio/kafka-go"
)

var TOPIC string
var PARTITION int
var KAFKA_ADDR string

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
