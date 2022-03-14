package redis

import (
	"ezpz/config"
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
	"strconv"
)

func NewClient() *redis.Client {
	db, err := strconv.Atoi(config.AppConfig()["redis_db"])
	if err != nil {
		log.Println(err)
	}

	return redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%v:%v", config.AppConfig()["redis_host"], config.AppConfig()["redis_port"]),
		Password: config.AppConfig()["redis_password"],
		DB:       db,
	})
}
