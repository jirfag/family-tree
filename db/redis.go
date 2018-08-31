package db

import (
	t "github.com/fredliang44/family-tree/graphql/types"
	"github.com/fredliang44/family-tree/utils"
	"log"

	"github.com/go-redis/redis"
)

// RedisClient locate mongo db client
var RedisClient = redisClient()

func redisClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     utils.AppConfig.Redis.Host + ":" + utils.AppConfig.Redis.Port,
		Password: utils.AppConfig.Redis.Password,
		DB:       utils.AppConfig.Redis.DB,
	})
	return client
}

// LoadUserCache is a func to load user data to redis
func LoadUserCache(user t.User) {
	//cache, _ := msgpack.Marshal(&user)
	CacheClient.Set(user.Username, user, -1)
}

// FetchUserCache is a func to fetch user data from redis
func FetchUserCache(username string) (user t.User, err error) {
	var res t.User
	raw, isExist := CacheClient.Get(username)
	if !isExist {
		log.Println("Get User from cache: ", err)
		res, err = FetchUserFromMongo(username)
		return res, err
	}

	return raw.(t.User), nil
}
