package db

import (
	"context"

	"github.com/go-redis/redis/v8"
)

func InitRedis(addr, password string, db int) (*redis.Client, error) {
	redisC := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
		PoolSize: 100, // 连接池大小
	})
	// 测试连接
	if _, err := redisC.Ping(context.Background()).Result(); err != nil {
		return nil, err
	}
	return redisC, nil
}
