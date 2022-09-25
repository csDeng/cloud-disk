package logic

import (
	"context"

	"cloud_disk/app/file/cmd/api/internal/svc"
	"cloud_disk/app/file/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserFileDelLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserFileDelLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFileDelLogic {
	return &UserFileDelLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserFileDelLogic) UserFileDel(req *types.UserFileDelRequest) (resp *types.UserFileDelResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
