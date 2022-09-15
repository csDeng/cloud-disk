package svc

import (
	"core/app/user/cmd/api/internal/config"
	"core/app/user/cmd/rpc/usercenter"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config

	UserRpcClient usercenter.UserCenter
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:        c,
		UserRpcClient: usercenter.NewUserCenter(zrpc.MustNewClient(c.UserRpcConf)),
	}
}
