package db

import (
	"family-tree/utils"
	"github.com/go-redis/redis"
)

var RedisClient = redisClient()

func redisClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     utils.AppConfig.Redis.Host + ":" + utils.AppConfig.Redis.Port,
		Password: utils.AppConfig.Redis.Password,
		DB:       utils.AppConfig.Redis.DB,
	})
	return client
}
