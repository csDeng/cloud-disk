package main

import (
	"flag"
	"fmt"

	"cloud_disk/app/common/vars"
	"cloud_disk/app/user/cmd/rpc/internal/config"
	"cloud_disk/app/user/cmd/rpc/internal/server"
	"cloud_disk/app/user/cmd/rpc/internal/svc"
	"cloud_disk/app/user/cmd/rpc/pb"
	"cloud_disk/app/user/cmd/rpc/user_helper"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/user.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	// 配置注入
	aesCfg := &vars.AesCfg{
		Secret: c.AesConf.Secret,
		IV:     c.AesConf.Iv,
	}

	tokenCfg := &vars.TokenConfig{
		TokenTime:        c.TokenConf.Token_time,
		RefreshTokenTime: c.TokenConf.Refresh_token_time,
		Secret:           c.TokenConf.Secret,
	}
	user_helper.InitCfg(aesCfg, tokenCfg)
	ctx := svc.NewServiceContext(c)
	svr := server.NewUserCenterServer(ctx)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		pb.RegisterUserCenterServer(grpcServer, svr)

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
