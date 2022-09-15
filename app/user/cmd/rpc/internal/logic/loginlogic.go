package logic

import (
	"context"

	"core/app/common/helper"
	"core/app/user/cmd/rpc/internal/svc"
	"core/app/user/cmd/rpc/pb"

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
	user, err := userModel.Login(Engine, in.Name, helper.Md5(in.Password))
	if err != nil {
		return nil, err
	}

	// 2. 生成 refresh_token
	refresh_token, err := helper.GenerateToken(user.Id, user.Identity, user.Name, true)
	if err != nil {
		return nil, err
	}
	// 3. 生成token
	token, err := helper.GenerateToken(user.Id, user.Identity, user.Name, false)
	if err != nil {
		return nil, err
	}
	return &pb.LoginResponse{
		Token:        token,
		RefreshToken: refresh_token,
	}, nil
}
