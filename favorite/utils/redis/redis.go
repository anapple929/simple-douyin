package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
)

var Ctx = context.Background()
var RdbUserVideo *redis.Client

// InitRedis 初始化Redis连接。
func InitRedis() {
	RdbUserVideo = redis.NewClient(&redis.Options{
		Addr:     "43.138.51.56:6379",
		Password: "292023",
		DB:       2, // 是否点赞信息存入 DB2.
	})
}
