package redisCon

import "github.com/go-redis/redis/v7"

var (
	RedisOp = &redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "toor",
		DB:       0,
		PoolSize: 10,
	}
)
