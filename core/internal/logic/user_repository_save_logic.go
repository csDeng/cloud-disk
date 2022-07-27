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

	ur := new(models.UserRepository)
	cnt, err := models.Engine.Where("user_identity = ?", userIdentity).And("repository_identity = ?", req.RepositoryIdentity).Count(ur)
	if err != nil {
		return nil, err
	}
	if cnt > 0 {
		return nil, errors.New("请勿重复关联")
	}

	ur = &models.UserRepository{
		Identity:           helper.GenerateUuid(),
		Name:               req.Name,
		Ext:                req.Ext,
		ParentId:           req.ParentId,
		UserIdentity:       userIdentity,
		RepositoryIdentity: req.RepositoryIdentity,
	}
	_, err = models.Engine.Insert(ur)
	if err != nil {
		return
	}
	resp = new(types.UserRepositoryResponse)
	resp.Identity = ur.Identity
	return
}
