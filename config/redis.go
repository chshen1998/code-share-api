package config

import (
	"log"

	"github.com/go-redis/redis"
)

var (
	RedisDB *redis.Client
)

func InitRedis() {
	RedisDB = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	pong, err := RedisDB.Ping().Result()
	if err != nil {
		log.Fatal(pong, err)
	}
}
