package rds

import (
	"core/app/common/vars"
	"fmt"

	"github.com/go-redis/redis/v8"
)

func GetRdsClient(c *vars.RedisConfig) *redis.Client {
	Redis := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", c.Server, c.Port),
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	return Redis
}