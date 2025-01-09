package producer

import (
	"context"
	"github.com/IBM/sarama"
	"liuhuig123demo/conf"
)

type HotelProducer struct {
	topic    string
	producer sarama.SyncProducer
}

func InitProducer(topic string) *HotelProducer {
	// todo: 填充kafka地址配置
	producer, err := sarama.NewSyncProducer([]string{}, conf.KafkaConf)
	if err != nil {
		panic(err)
	}
	return &HotelProducer{
		topic:    topic,
		producer: producer,
	}
}

func (c *HotelProducer) SendMessage(ctx context.Context, message string) error {
	msg := &sarama.ProducerMessage{}
	msg.Topic = c.topic
	msg.Value = sarama.StringEncoder(message)

	_, _, err := c.producer.SendMessage(msg)
	if err != nil {
		return err
	}
	return nil
}
