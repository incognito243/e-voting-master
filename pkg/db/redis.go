package db

import (
	"context"
	"errors"

	"e-voting-mater/configs"

	"github.com/redis/go-redis/v9"
)

func InitRedis() (redis.UniversalClient, error) {
	redisAddresses := configs.G.Redis.InitAddress
	if len(redisAddresses) == 0 {
		return nil, errors.New("redis host is empty")
	}

	redisClient := redis.NewUniversalClient(&redis.UniversalOptions{
		Addrs:      redisAddresses,
		MasterName: configs.G.Redis.MasterName,
		Password:   configs.G.Redis.Password,
	})
	if _, err := redisClient.Ping(context.Background()).Result(); err != nil {
		return nil, err
	}
	return redisClient, nil
}
