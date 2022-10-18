package logic

import (
	"context"
	"fmt"

	"cloud_disk/app/common/helper"
	"cloud_disk/app/common/vars"
	"cloud_disk/app/file/cmd/api/internal/svc"
	"cloud_disk/app/file/cmd/api/internal/types"
	"cloud_disk/app/file/models"

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
	Model, Engine := l.svcCtx.UserRepository, l.svcCtx.Engine
	has, err := Model.CheckParentIfExisted(Engine, userIdentity, req.ParentId)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, vars.ErrParentNotExisted
	}

	has, err = Model.CheckIfHasNameWithPId(Engine, req.ParentId, userIdentity, req.Name)
	if err != nil {
		return nil, err
	}
	if has {
		return nil, fmt.Errorf("【%d】 目录下存在 【%s】 文件", req.ParentId, req.Name)
	}
	data := &models.UserRepository{
		Identity:     helper.GenerateUuid(),
		ParentId:     req.ParentId,
		UserIdentity: userIdentity,
		Name:         req.Name,
	}
	ok, err := Model.CreateFolder(Engine, data)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, vars.ErrCreateFolder
	}
	resp = new(types.UserFolderCreateResponse)
	resp.Identity = data.Identity
	return resp, nil
}
