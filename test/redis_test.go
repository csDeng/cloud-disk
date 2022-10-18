package test

import (
	"cloud_disk/app/common/vars"
	"context"
	"fmt"
	"log"
	"testing"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func TestRedis(t *testing.T) {
	config := vars.RedisConfig{
		Server:      "",
		Port:        6380,
		RedisPrefix: "test",
		Password:    "",
	}
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", config.Server, config.Port),
		Password: config.Password, // no password set
		DB:       0,               // use default DB
	})
	err := rdb.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		panic(err)
	}
	script := `
	-- old token key , new token key, Expire
	local Expire = tonumber(ARGV[3])
	if(not KEYS[3]) then
		return false
	elseif Expire == nil then
		return false	
	end
	redis.call("SET",KEYS[1],ARGV[1],"EX",ARGV[3])
	redis.call("SET",KEYS[2],ARGV[2],"EX",ARGV[3])
	return true
	`
	// res, err := rdb.Eval(ctx, script, []string{"key1", "key2", "k"}, "ag1", "arg2", 1000).Result()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Printf("Eval T=%T v=%v", res, res)

	v, err := rdb.ScriptLoad(ctx, script).Result()
	fmt.Printf("%+v %+v", v, err)

	err = rdb.EvalSha(ctx, "ce44b37a62034ec65136828fb551bcbc890a58b2", []string{"k1", "k2", "k3"}, 1, 2).Err()
	if err != nil {
		log.Fatal(err)
	}
}
