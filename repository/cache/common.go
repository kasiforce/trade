package cache

import (
	"context"
	"fmt"
	conf "github.com/kasiforce/trade/config"
	"github.com/kasiforce/trade/pkg/util"
	"github.com/redis/go-redis/v9"
)

// RedisClient Redis缓存客户端单例
var RedisClient *redis.Client
var RedisContext = context.Background()

// InitCache 在中间件中初始化redis链接
func InitCache() {
	rConfig := conf.Config.Redis
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", rConfig.RedisHost, rConfig.RedisPort),
		Username: rConfig.RedisUsername,
		Password: rConfig.RedisPassword,
		DB:       rConfig.RedisDbName,
	})
	_, err := client.Ping(RedisContext).Result()
	if err != nil {
		util.LogrusObj.Info(err)
		panic(err)
	}
	RedisClient = client
}
