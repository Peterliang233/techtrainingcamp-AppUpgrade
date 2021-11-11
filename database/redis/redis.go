package redis

import (
	"context"
	"fmt"
	"log"

	"github.com/Peterliang233/techtrainingcamp-AppUpgrade/config"

	"github.com/go-redis/redis/v8"
)

var RedisClient *redis.Client

// InitRedis 初始化redis服务器
func InitRedis() {
	redisAddr := fmt.Sprintf("%s:%s",
		config.RedisSetting.RdHost,
		config.RedisSetting.RdPort)

	RedisClient = redis.NewClient(
		&redis.Options{
			Addr:     redisAddr,
			Password: config.RedisSetting.RbPassword,
			DB:       0,
		},
	)
	pong, err := RedisClient.Ping(context.Background()).Result()
	if err != nil {
		log.Fatalf("Redis 启动错误 %v %v\n", pong, err)
	}
	log.Println("redis start success")
}
