package contact

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
)

var RedisClient *redis.Client

func InitRedis() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", Config.Redis.Host, Config.Redis.Port),
		Password: Config.Redis.Password, // no password set
		DB:       Config.Redis.Db,       // use default DB
	})

	_, err := RedisClient.Ping(context.Background()).Result()

	if err != nil {
		log.Fatalln(err)
	}
}

func RedisClose() error {
	return RedisClient.Close()
}
