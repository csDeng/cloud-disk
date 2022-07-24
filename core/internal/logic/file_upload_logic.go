package logic

import (
	"context"

	"core/core/helper"
	"core/core/internal/svc"
	"core/core/internal/types"
	"core/models"

	"github.com/zeromicro/go-zero/core/logx"
)

type FileUploadLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFileUploadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FileUploadLogic {
	return &FileUploadLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FileUploadLogic) FileUpload(req *types.FileUploadRequest) (resp *types.FileUploadResponse, err error) {
	rp := &models.RepositoryPool{
		Identity: helper.GenerateUuid(),
		Hash:     req.Hash,
		Name:     req.Name,
		Ext:      req.Ext,
		Size:     int(req.Size),
		Path:     req.Path,
	}

	_, err = models.Engine.Insert(rp)
	if err != nil {
		return nil, err
	}
	resp = &types.FileUploadResponse{
		Identity: rp.Identity,
		Name:     rp.Name,
		Ext:      rp.Ext,
		Size:     int64(rp.Size),
		Path:     rp.Path,
	}
	return resp, nil
}
