package logic

import (
	"context"
	"errors"

	"core/core/helper"
	"core/core/internal/svc"
	"core/core/internal/types"
	"core/models"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserFolderCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserFolderCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFolderCreateLogic {
	return &UserFolderCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserFolderCreateLogic) UserFolderCreate(req *types.UserFolderCreateRequest, userIdentity string) (resp *types.UserFolderCreateResponse, err error) {
	cnt, err := models.Engine.Where("parent_id = ? AND name = ?", req.ParentId, req.Name).Count(new(models.UserRepository))
	if err != nil {
		return nil, err
	}
	if cnt >= 1 {
		return nil, errors.New("当前文件夹已存在该名称")
	}

	// 创建文件夹
	data := &models.UserRepository{
		Identity:     helper.GenerateUuid(),
		ParentId:     req.ParentId,
		UserIdentity: userIdentity,
		Name:         req.Name,
	}
	_, err = models.Engine.Insert(data)
	if err != nil {
		return nil, err
	}
	return &types.UserFolderCreateResponse{
		Identity: data.Identity,
	}, nil

}
