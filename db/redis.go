package db

import (
	"family-tree/utils"
	"fmt"
	"github.com/go-redis/redis"
)

func Client() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     utils.AppConfig.Redis.Host + ":" + utils.AppConfig.Redis.Port,
		Password: utils.AppConfig.Redis.Password, // no password set
		DB:       utils.AppConfig.Redis.DB,       // use default DB
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)

	return client
}

var RedisClient = Client()
