package conf

import "github.com/IBM/sarama"

var KafkaConf *sarama.Config

func init() {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	KafkaConf = config
}
