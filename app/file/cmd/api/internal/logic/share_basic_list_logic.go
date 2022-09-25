package logic

import (
	"context"

	"cloud_disk/app/file/cmd/api/internal/svc"
	"cloud_disk/app/file/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ShareBasicListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewShareBasicListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShareBasicListLogic {
	return &ShareBasicListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ShareBasicListLogic) ShareBasicList(req *types.ShareBasicListRequest) (resp *types.ShareBasicListResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
