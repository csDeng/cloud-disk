package logic

import (
	"context"

	"cloud_disk/app/user/cmd/rpc/internal/svc"
	"cloud_disk/app/user/cmd/rpc/pb"
	"cloud_disk/app/user/cmd/rpc/user_helper"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetIdentityWithTokenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetIdentityWithTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetIdentityWithTokenLogic {
	return &GetIdentityWithTokenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetIdentityWithTokenLogic) GetIdentityWithToken(in *pb.GetIdentityWithTokenRequest) (*pb.GetIdentityWithTokenResponse, error) {
	uc, err := user_helper.ParseToken(in.Token)
	if err != nil {
		return nil, err
	}
	return &pb.GetIdentityWithTokenResponse{
		Identity: uc.Identity,
	}, nil
}
