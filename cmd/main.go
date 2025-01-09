package cmd

import (
	"context"
	"encoding/json"
	"flag"
	"io/ioutil"
	"liuhuig123demo/cmd/consumer"
	"liuhuig123demo/cmd/producer"
	"liuhuig123demo/conf"
	"liuhuig123demo/internal/dto"
	"strconv"
)

const (
	scriptNameConsumer = "consumer"
	scriptNameProducer = "producer"
)

var (
	startFlag string
	workers   int
	queueName string
	data      string
)

func init() {
	flag.StringVar(&startFlag, "start", "", "")
	flag.IntVar(&workers, "workers", 0, "")
	flag.StringVar(&queueName, "queue", "", "")
	flag.StringVar(&data, "data", "", "")
}

func main() {
	flag.Parse()

	switch startFlag {
	case scriptNameConsumer:
		hotelConsumer := consumer.InitConsumer(workers, queueName)
		hotelConsumer.Start()
	case scriptNameProducer:
		hotelProducer := producer.InitProducer(queueName)

		ctx := context.Background()
		fileData, err := ioutil.ReadFile(data)
		if err != nil {
			return
		}

		var tasksJson dto.TasksJson
		err = json.Unmarshal(fileData, &tasksJson)
		if err != nil {
			return
		}

		for _, task := range tasksJson.Tasks {
			taskId, err := conf.RedisCli.Incr(ctx, "hotel_task_id").Result()
			if err != nil {
				return
			}
			task.TaskId = strconv.FormatInt(taskId, 10)

			taskMsg, err := json.Marshal(task)
			if err != nil {
				return
			}

			_ = hotelProducer.SendMessage(ctx, string(taskMsg))
		}
	default:
		return
	}

}
