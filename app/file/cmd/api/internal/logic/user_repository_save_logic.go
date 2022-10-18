package logic

import (
	"context"

	"cloud_disk/app/common/helper"
	"cloud_disk/app/common/vars"
	"cloud_disk/app/file/cmd/api/internal/svc"
	"cloud_disk/app/file/cmd/api/internal/types"
	"cloud_disk/app/file/models"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserRepositorySaveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserRepositorySaveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRepositorySaveLogic {
	return &UserRepositorySaveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserRepositorySaveLogic) UserRepositorySave(req *types.UserRepositoryRequest, userIdentity string) (resp *types.UserRepositoryResponse, err error) {
	// 1. 检查文件是否存在
	Engine := l.svcCtx.Engine
	UserRepositoryModel := l.svcCtx.UserRepository
	PoolModel := l.svcCtx.PoolModel
	file, err := PoolModel.GetFileWithIdentity(Engine, req.RepositoryIdentity)
	if err != nil {
		return nil, err
	}
	if len(file.Identity) == 0 {
		return nil, vars.ErrFileNotExisted
	}
	// 2. 检查当前用户是否与该文件关联

	up, err := UserRepositoryModel.CheckUserIfHasFile(Engine, userIdentity, req.RepositoryIdentity)
	if err != nil {
		return nil, err
	}

	if len(up.Identity) > 0 {
		return nil, vars.ErrFileHasExisted
	}

	// 3. 关联文件
	up = &models.UserRepository{
		Identity:           helper.GenerateUuid(),
		ParentId:           req.ParentId,
		UserIdentity:       userIdentity,
		RepositoryIdentity: file.Identity,
		Ext:                file.Ext,
		Name:               file.Name,
	}
	up, err = UserRepositoryModel.AddFile(Engine, up)
	if err != nil {
		return nil, err
	}
	resp = &types.UserRepositoryResponse{
		Identity: up.Identity,
	}
	return resp, nil
}
