package logic

import (
	"context"

	"core/core/helper"
	"core/core/internal/svc"
	"core/core/internal/types"

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
	err = helper.ChunkSuccess(req)
	resp = new(types.FileChunkSuccessResponse)
	return
}
