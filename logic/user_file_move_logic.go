package logic

import (
	"context"

	"cloud_disk/app/internal/svc"
	"cloud_disk/app/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserFileMoveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserFileMoveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFileMoveLogic {
	return &UserFileMoveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserFileMoveLogic) UserFileMove(req *types.UserFileMoveRequest, userIdentity string) (resp *types.UserFileMoveResponse, err error) {
	// engine := models.Engine
	// // 判断文件是否存在
	// has, err := engine.Where("identity = ? AND user_identity = ?", req.Identity, userIdentity).Exist(new(models.UserRepository))
	// if err != nil {
	// 	return nil, err
	// }
	// if !has {
	// 	return nil, fmt.Errorf("identity = %s 's file is not existed", req.Identity)
	// }
	// // 判断 parent_id 是否存在
	// has, err = engine.Where("id = ? AND user_identity = ? AND LENGTH(repository_identity) = 0", req.ParentId, userIdentity).Exist(new(models.UserRepository))
	// if err != nil {
	// 	return nil, err
	// }
	// if !has {
	// 	return nil, fmt.Errorf("parent_id = %d 's folder is not existed", req.ParentId)
	// }
	// // 判断是否移动成功
	// ok, err := engine.Where("identity = ? ", req.Identity).Update(&models.UserRepository{ParentId: req.ParentId})
	// if err != nil {
	// 	return nil, err
	// }
	// if ok < 1 {
	// 	return nil, errors.New("文件移动失败")
	// }
	return
}
