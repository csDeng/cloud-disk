package config

import (
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	Mode string
	zrpc.RpcServerConf

	EmailRpcConf zrpc.RpcClientConf

	Mysql struct {
		Server   string
		Port     int
		Db       string
		User     string
		Password string
	}

	RedisConf struct {
		Server string
		Port   int
		Prefix string
	}

	RandCodeLength int

	RandCodeExpire int

	TokenConf struct {
		Token_time         int
		Refresh_token_time int
		Secret             string
	}

	AesConf struct {
		Secret string
		Iv     string
	}
}
