package db

import (
	"context"
	"github.com/redis/go-redis/v9"
)

var RDB *redis.Client

func init() {
	RDB = redis.NewClient(&redis.Options{
		Addr:     "112.126.71.240:6379",
		Password: "dockerredis", // 没有密码，默认值
		DB:       0,             // 默认DB 0
	})
	ctx := context.Background()
	_, err := RDB.Ping(ctx).Result()
	if err != nil {
		panic(err)
	}
}
