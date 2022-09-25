package logic

import (
	"context"

	"cloud_disk/app/user/cmd/api/internal/svc"
	"cloud_disk/app/user/cmd/api/internal/types"
	"cloud_disk/app/user/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLoginLogic {
	return &UserLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserLoginLogic) UserLogin(req *types.LoginRequest) (resp *types.LoginResponse, err error) {
	// todo: add your logic here and delete this line
	rpcResp, err := l.svcCtx.UserRpcClient.Login(l.ctx, &pb.LoginRequest{
		Name:     req.Name,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}
	resp = new(types.LoginResponse)
	resp.Token = rpcResp.Token
	resp.RefreshToken = rpcResp.RefreshToken
	return resp, nil
}
