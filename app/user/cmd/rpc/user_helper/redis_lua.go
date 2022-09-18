package user_helper

import (
	"context"

	"github.com/go-redis/redis/v8"
)

// 输入参数
// old_refresh_token val
// new_refresh_token_key val
// time ExpireTime
func getTokenScript() string {
	lua := `
	-- old token key , new token key, Expire
	local Expire = tonumber(ARGV[3])
	if(not KEYS[3]) then
		return 2
	elseif Expire == nil then
		return 2	
	end
	redis.call("SET",KEYS[1],ARGV[1],"EX",ARGV[3])
	redis.call("SET",KEYS[2],ARGV[2],"EX",ARGV[3])
	return 1
	`
	return lua
}

var sha string

func GetSha(ctx context.Context, rdb *redis.Client) (string, error) {
	if len(sha) != 0 {
		return sha, nil
	}
	sha, err := rdb.ScriptLoad(ctx, getTokenScript()).Result()
	if err != nil {
		return "", err
	}
	return sha, nil
}
