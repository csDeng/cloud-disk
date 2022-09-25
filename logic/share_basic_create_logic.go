package logic

import (
	"context"

	"cloud_disk/app/internal/svc"
	"cloud_disk/app/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ShareBasicCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewShareBasicCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShareBasicCreateLogic {
	return &ShareBasicCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ShareBasicCreateLogic) ShareBasicCreate(req *types.ShareBasicCreateRequest, userIdentity string) (resp *types.ShareBasicCreateResponse, err error) {
	// engine := models.Engine
	// // 判断文件是否存在
	// has, err := engine.Where("repository_identity = ? AND user_identity = ?", req.RepositoryIdentity, userIdentity).Exist(new(models.UserRepository))
	// if err != nil {
	// 	return nil, err
	// }
	// if !has {
	// 	return nil, fmt.Errorf("repository_identity = %s 's file is not existed", req.RepositoryIdentity)
	// }
	// data := &models.ShareBasic{
	// 	Identity:           helper.GenerateUuid(),
	// 	UserIdentity:       userIdentity,
	// 	RepositoryIdentity: req.RepositoryIdentity,
	// 	ExpiredTime:        req.ExpiredTime,
	// 	ClickNum:           0,
	// }
	// _, err = engine.Insert(data)
	// if err != nil {
	// 	return nil, err
	// }
	// resp = new(types.ShareBasicCreateResponse)
	// resp.Identity = data.Identity
	return
}
