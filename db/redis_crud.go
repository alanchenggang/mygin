package db

import (
	"context"
	"github.com/redis/go-redis/v9"
	"mygin/log"
	"strconv"
	"time"
)

var contect = context.Background()

func CRUD_Redis() {
	for i := 0; i < 100; i++ {
		RedisSet("name"+strconv.Itoa(i), "coder4j")
	}
}

// RedisSet 设置redis
func RedisSet(key string, value interface{}) {
	RDB.Set(contect, key, value, 50*time.Second)
}

// HandleRedisErr 处理redis错误
func HandleRedisErr(err error) {
	if err != nil && err == redis.Nil {
		log.Log.Warnln("key does not exist")
	} else if err != nil {
		log.Log.Warnln(err.Error())
	}
}
