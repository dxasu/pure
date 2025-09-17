package db

import (
	"github.com/segmentio/kafka-go"
)

var KafkaWriter *kafka.Writer

func InitKafka(brokers []string, topic string) {
	KafkaWriter = &kafka.Writer{
		Addr:     kafka.TCP(brokers...),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	}
}

func CloseKafka() {
	if KafkaWriter != nil {
		KafkaWriter.Close()
	}
}

// brokers := []string{"127.0.0.1:9092"}
// config := sarama.NewConfig()
// config.Net.SASL.Enable = true
// config.Net.SASL.User = "root"
// config.Net.SASL.Password = "123456"
// producer, err := sarama.NewSyncProducer(brokers, config)
