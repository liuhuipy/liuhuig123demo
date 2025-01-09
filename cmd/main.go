package cmd

import "flag"

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

	case scriptNameProducer:

	default:
		return
	}

}
