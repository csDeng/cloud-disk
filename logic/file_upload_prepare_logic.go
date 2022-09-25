package logic

import (
	"context"

	"cloud_disk/app/internal/svc"
	"cloud_disk/app/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FileUploadPrepareLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFileUploadPrepareLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FileUploadPrepareLogic {
	return &FileUploadPrepareLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FileUploadPrepareLogic) FileUploadPrepare(req *types.FileUploadPrepareRequest) (resp *types.FileUploadPrepareResponse, err error) {
	// p := new(models.RepositoryPool)
	// engine := models.Engine
	// has, err := engine.Where("hash = ?", req.Md5).Get(p)
	// if err != nil {
	// 	return nil, fmt.Errorf("hash = %s 的文件已存在", req.Md5)
	// }
	// resp = new(types.FileUploadPrepareResponse)
	// if has {
	// 	// 如果文件已存在，秒传
	// 	resp.Identity = p.Identity
	// } else {
	// 	// 获取UploadId 进行分片上传
	// 	key, id, err := helper.ChunkInit(req.Ext)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	resp.Key = key
	// 	resp.UploadId = id
	// }
	return
}
