package logic

import (
	"context"
	"errors"

	"core/core/internal/svc"
	"core/core/internal/types"
	"core/models"

	"github.com/zeromicro/go-zero/core/logx"
)

type EditFileNameLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewEditFileNameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EditFileNameLogic {
	return &EditFileNameLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *EditFileNameLogic) EditFileName(req *types.EditFileNameRequest, userIdentity string) (resp *types.EditFileNameResponse, err error) {
	data := &models.UserRepository{Name: req.Name}
	cnt, err := models.Engine.Where("identity = ?", req.Identity).Count(new(models.UserRepository))
	if err != nil {
		return nil, err
	}

	if cnt < 1 {
		return nil, errors.New("文件不存在, 请检查 identity")
	}
	cnt, err = models.Engine.Where("parent_id = (SELECT parent_id FROM user_repository ur WHERE ur.identity = ?)  AND name = ?", req.Identity, req.Name).Count(new(models.UserRepository))
	if err != nil {
		return nil, err
	}
	if cnt == 1 {
		return nil, errors.New("该文件夹下已存在该名称, 请输入新的文件名")
	}
	update, err := models.Engine.Where("identity = ?", req.Identity).Update(data)
	if err != nil {
		return nil, err
	}
	if update < 1 {
		return nil, errors.New("没有发生变化")
	}
	return
}
