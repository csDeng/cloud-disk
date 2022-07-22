package test

import (
	"context"
	"core/core/helper"
	"fmt"
	"testing"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func TestRedis(t *testing.T) {
	config := helper.RedisConfigObject
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", config.Server, config.Port),
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	v, err := rdb.Exists(ctx, "key").Result()

	// v, err := rdb.Do(ctx, "TTL", "key").Result()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%T, %v", v, v)

	// err := rdb.Set(ctx, "key", "value", 0).Err()
	// if err != nil {
	// 	panic(err)
	// }

	// val, err := rdb.Get(ctx, "key").Result()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("key", val)

	// val2, err := rdb.Get(ctx, "key2").Result()
	// if err == redis.Nil {
	// 	fmt.Println("key2 does not exist")
	// } else if err != nil {
	// 	panic(err)
	// } else {
	// 	fmt.Println("key2", val2)
	// }
	// Output: key value
	// key2 does not exist
}
