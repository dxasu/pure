package db

import (
	"context"

	"github.com/go-redis/redis/v8"
)

var RedisClient *redis.Client

func InitRedis(addr, password string, db int) error {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
		PoolSize: 100, // 连接池大小
	})
	// 测试连接
	if _, err := RedisClient.Ping(context.Background()).Result(); err != nil {
		return err
	}
	return nil
}
