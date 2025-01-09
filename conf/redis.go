package conf

import (
	"context"
	"github.com/go-redis/redis/v8"
)

var RedisCli *redis.Client

func init() {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1",
		Username: "",
		Password: "",
	})
	_, err := redisClient.Ping(context.Background()).Result()
	if err != nil {
		panic(err)
	}
	RedisCli = redisClient
}
