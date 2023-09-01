package service

import (
	"CRUD/pkg/model"
	"encoding/json"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type KafkaSender struct {
	producerChan  chan kafka.Event
	kafkaProducer *kafka.Producer
}

func NewKafkaSender(producerChan chan kafka.Event, kafkaProducer *kafka.Producer) *KafkaSender {
	return &KafkaSender{producerChan: producerChan, kafkaProducer: kafkaProducer}
}

func (c *KafkaSender) SentMessage(message *logrus.Entry) {
	msg, _ := json.Marshal(model.KafkaMessage{
		Level:    message.Level.String(),
		Msg:      message.Message,
		TimeSent: message.Time,
	})

	topic := viper.GetString("kafka.topics.producer")
	c.kafkaProducer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &topic,
			Partition: kafka.PartitionAny},
		Value: msg,
	},
		c.producerChan)

	e := <-c.producerChan
	m := e.(*kafka.Message)

	if m.TopicPartition.Error != nil {
		fmt.Printf("Delivery failed: %v\n", m.TopicPartition.Error)
	} else {
		fmt.Printf("Delivered message to topic %s [%d] at offset %v\n",
			*m.TopicPartition.Topic, m.TopicPartition.Partition, m.TopicPartition.Offset)
	}
}
