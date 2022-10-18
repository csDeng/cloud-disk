package logic

import (
	"context"

	"cloud_disk/app/common/vars"
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

func (l *UserFileDelLogic) UserFileDel(req *types.UserFileDelRequest, userIdentity string) (resp *types.UserFileDelResponse, err error) {
	Engine, UserRep := l.svcCtx.Engine, l.svcCtx.UserRepository

	// 检查是否是一个目录
	b, err := UserRep.CheckFileIsFolder(Engine, userIdentity, req.Identity)
	if err != nil {
		return nil, err
	}
	if b {
		return nil, vars.ErrFileIsFolder
	}

	b, err = UserRep.DelUserFile(Engine, req.Identity, userIdentity)
	if err != nil {
		return nil, err
	}
	if !b {
		return nil, vars.ErrFileNotExisted
	}
	return nil, nil
}
