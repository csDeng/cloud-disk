package logic

import (
	"context"
	"time"

	"cloud_disk/app/common/helper"
	"cloud_disk/app/common/vars"
	"cloud_disk/app/user/cmd/rpc/internal/svc"
	"cloud_disk/app/user/cmd/rpc/pb"
	"cloud_disk/app/user/cmd/rpc/user_helper"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *pb.LoginRequest) (*pb.LoginResponse, error) {
	userModel, Engine := l.svcCtx.UserModel, l.svcCtx.Engine
	rds := l.svcCtx.RdsCli
	user, err := userModel.Login(Engine, in.Name, helper.Md5(in.Password))
	if err != nil {
		return nil, err
	}

	// 2. 生成 refresh_token

	refresh_token, err := user_helper.GenerateToken(user.Id, user.Identity, user.Name, true)
	if err != nil {
		return nil, err
	}
	// 3. 生成token
	token, err := user_helper.GenerateToken(user.Id, user.Identity, user.Name, false)
	if err != nil {
		return nil, err
	}

	prefix := l.svcCtx.Config.RedisConf.Prefix
	rtk := helper.GetRefreshTokenKey(prefix, refresh_token)

	err = rds.Set(l.ctx, rtk, vars.TOKEN_STORE,
		time.Minute*time.Duration(l.svcCtx.Config.TokenConf.Refresh_token_time)).Err()
	if err != nil {
		return nil, err
	}
	return &pb.LoginResponse{
		Token:        token,
		RefreshToken: refresh_token,
		Identity:     user.Identity,
	}, nil
}
