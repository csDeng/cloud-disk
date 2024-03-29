package logic

import (
	"context"

	"cloud_disk/app/file/cmd/api/internal/svc"
	"cloud_disk/app/file/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FileChunkUploadLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFileChunkUploadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FileChunkUploadLogic {
	return &FileChunkUploadLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FileChunkUploadLogic) FileChunkUpload(req *types.FileChunkUploadRequest) (resp *types.FileChunkUploadResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
