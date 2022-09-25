package logic

import (
	"context"

	"cloud_disk/app/internal/svc"
	"cloud_disk/app/internal/types"

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
	// todo: add your logic here and delete this line
	// _, err = models.Engine.Where("identity = ? AND user_identity = ?", req.Identity, userIdentity).Delete(new(models.UserRepository))

	// if err != nil {
	// 	return nil, err
	// }
	return
}
