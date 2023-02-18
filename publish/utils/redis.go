package utils

import (
	"context"
	"github.com/go-redis/redis/v8"
)

var Ctx = context.Background()
var RdbVideoId *redis.Client

// InitRedis 初始化Redis连接。
func InitRedis() {

	RdbVideoId = redis.NewClient(&redis.Options{
		Addr:     "43.138.51.56:6379",
		Password: "292023",
		DB:       1, // 视频信息存入 DB1.
	})
}
