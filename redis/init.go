package redis

import (
	"core/core/helper"
	"fmt"

	"github.com/go-redis/redis/v8"
)

var Redis *redis.Client

func init() {
	config := helper.RedisConfigObject
	Redis = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", config.Server, config.Port),
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}
