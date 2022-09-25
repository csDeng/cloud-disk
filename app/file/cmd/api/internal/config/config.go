package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf

	UserRpcConf zrpc.RpcClientConf

	Mysql struct {
		Server   string
		Port     int
		Db       string
		User     string
		Password string
	}
}
