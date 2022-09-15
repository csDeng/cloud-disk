package config

import "github.com/zeromicro/go-zero/zrpc"

type EmailConfigType struct {
	Server   string
	Port     int
	From     string
	Password string
	Second   int
}

type Config struct {
	zrpc.RpcServerConf
	EmailConfig EmailConfigType
}
