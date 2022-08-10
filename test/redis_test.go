package test

import (
	"context"
	"core/core/helper"
	"fmt"
	"strconv"
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
	// v, err := rdb.Exists(ctx, "key").Result()

	// // v, err := rdb.Do(ctx, "TTL", "key").Result()
	// if err != nil {
	// 	t.Fatal(err)
	// }
	// t.Logf("%T, %v", v, v)

	// err := rdb.Set(ctx, "key", "value", time.Second*30).Err()
	// if err != nil {
	// 	panic(err)
	// }

	val, err := rdb.Get(ctx, "cloud_disk:refresh_token:1eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJZCI6NiwiSWRlbnRpdHkiOiJjMjc4ZWRjNi0wYjJkLTQ3NzMtOWQ0Yy0xMzIzZDFhMjZjMjQiLCJOYW1lIjoibXluYW1lIiwiUmVmcmVzaFRva2VuSWQiOiI4OTQ5YjYxNS00N2UwLTRjM2QtYWY0My0zMjQyZWYzYzFlNzMiLCJleHAiOjE2NjAxNDE1OTl9.0WDuIORD_0qxwxhwLRxNhtoMaXfLmJyLtmlldnVDdMw").Result()
	if err != nil && err != redis.Nil {
		t.Log(err)
	}
	fmt.Println("key", val)
	v, err := strconv.Atoi(val)
	if err != nil {
		t.Fail()
	}
	t.Log(v == 1)

	// val, err := rdb.Do(ctx, "EXISTS", "cloud_disk:1refresh_token:eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJZCI6NiwiSWRlbnRpdHkiOiJjMjc4ZWRjNi0wYjJkLTQ3NzMtOWQ0Yy0xMzIzZDFhMjZjMjQiLCJOYW1lIjoibXluYW1lIiwiUmVmcmVzaFRva2VuSWQiOiI4OTQ5YjYxNS00N2UwLTRjM2QtYWY0My0zMjQyZWYzYzFlNzMiLCJleHAiOjE2NjAxNDE1OTl9.0WDuIORD_0qxwxhwLRxNhtoMaXfLmJyLtmlldnVDdMw").Result()
	// if err != nil {
	// 	t.Fatal(err)
	// }
	// fmt.Printf("%T %v", val, val)
	// t.Log(val == int64(1))

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
