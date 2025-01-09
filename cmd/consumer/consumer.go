package consumer

import (
	"context"
	"encoding/json"
	"github.com/IBM/sarama"
	"github.com/panjf2000/ants/v2"
	"liuhuig123demo/internal/dto"
	"liuhuig123demo/internal/service"
)

type HotelConsumer struct {
	topic    string
	pool     *ants.Pool
	consumer sarama.Consumer
}

func InitConsumer(workers int, topic string) *HotelConsumer {
	pool, err := ants.NewPool(workers)
	if err != nil {
		panic(err)
	}

	client, err := sarama.NewClient([]string{}, nil)
	if err != nil {
		panic(err)
	}

	consumer, err := sarama.NewConsumerFromClient(client)
	if err != nil {
		panic(err)
	}

	return &HotelConsumer{topic: topic, pool: pool, consumer: consumer}

}

func (c *HotelConsumer) consumeKafkaPartition(consumer sarama.PartitionConsumer) {
	for msg := range consumer.Messages() {
		_ = c.pool.Submit(func() {
			var task dto.Task
			err := json.Unmarshal(msg.Value, &task)
			if err != nil {
				return
			}

			_ = service.HotelService.HandleTask(context.Background(), &task)
		})
	}
}

func (c *HotelConsumer) Start() {
	partitions, err := c.consumer.Partitions(c.topic)
	if err != nil {
		panic(err)
	}

	for _, partition := range partitions {
		p, err := c.consumer.ConsumePartition(c.topic, partition, sarama.OffsetNewest)
		if err != nil {
			panic(err)
		}

		go c.consumeKafkaPartition(p)
	}

}

func (c *HotelConsumer) Stop() {
	// todo
	err := c.consumer.Close()
	if err != nil {
		return
	}
}
