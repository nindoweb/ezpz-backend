package redis

import (
	"fmt"

	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
)

func NewClient() *redis.Client {

	return redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%v:%v", viper.GetString("REDIS_HOST"), viper.GetString("REDIS_PORT")),
		Password: viper.GetString("REDIS_PASSWORD"),
		DB: viper.GetInt("REDIS_NAME"),
	})
}
