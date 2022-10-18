package logic

import (
	"context"
	"errors"
	"fmt"

	"cloud_disk/app/common/vars"
	"cloud_disk/app/file/cmd/api/internal/svc"
	"cloud_disk/app/file/cmd/api/internal/types"

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
	UPModel := l.svcCtx.UserRepository
	Engine := l.svcCtx.Engine
	user_repository, err := UPModel.CheckIfExistedIdentity(Engine, userIdentity, req.Identity)
	if err != nil {
		return nil, err
	}
	if len(user_repository.Identity) == 0 {
		return nil, vars.ErrFileNotExisted
	}

	has, err := UPModel.CheckIfHasNameWithPId(Engine, user_repository.ParentId, userIdentity, req.Name)
	if err != nil {
		return nil, err
	}
	if has {
		return nil, fmt.Errorf("%d 目录下已存在名称： %s", user_repository.ParentId, req.Name)
	}
	ok, err := UPModel.EditFileName(Engine, userIdentity, req.Identity, req.Name)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, errors.New("文件名修改失败")
	}
	resp = new(types.EditFileNameResponse)
	return resp, nil
}
