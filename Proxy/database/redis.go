package database

import (
	"context"

	"github.com/redis/go-redis/v9"
	"github.com/vesperismc/vesperis/proxy/logger"
)

func initializeRedis() {
	logger := logger.GetLogger()
	opt, err := redis.ParseURL(GetRedisUrl())
	if err != nil {
		logger.Error("Error parsing Redis URL: " + err.Error())
		panic(err)
	}

	client := redis.NewClient(opt)

	_, err = client.Ping(context.Background()).Result()
	if err != nil {
		logger.Error("Error connecting to the Redis Database: " + err.Error())
		panic(err)
	}

	logger.Info("Successfully initialized the Redis Database.")
}
