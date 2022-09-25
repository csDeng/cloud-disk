package logic

import (
	"context"

	"cloud_disk/app/file/cmd/api/internal/svc"
	"cloud_disk/app/file/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FileChunkSuccessLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFileChunkSuccessLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FileChunkSuccessLogic {
	return &FileChunkSuccessLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FileChunkSuccessLogic) FileChunkSuccess(req *types.FileChunkSuccessRequest) (resp *types.FileChunkSuccessResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
