package logic

import (
	"context"

	"cloud_disk/app/common/vars"
	"cloud_disk/app/file/cmd/api/internal/svc"
	"cloud_disk/app/file/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserFileListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserFileListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFileListLogic {
	return &UserFileListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserFileListLogic) UserFileList(req *types.UserFileListRequest, userIdentity string) (resp *types.UserFileListResponse, err error) {
	Engine, UserRep := l.svcCtx.Engine, l.svcCtx.UserRepository
	resp = new(types.UserFileListResponse)
	cnt, err := UserRep.GetTotalWithParentId(Engine, userIdentity, req.Id)
	if err != nil {
		return nil, err
	}
	if cnt <= 0 {
		return resp, nil
	}
	size := req.Size
	if size == 0 {
		size = vars.PageSize
	}
	page := req.Page
	if page <= 0 {
		page = 1
	}
	offset := (page - 1) * size
	list, err := UserRep.GetUserFileList(Engine, userIdentity, req.Id, size, offset)
	if err != nil {
		return nil, err
	}
	res := make([]*types.UserFile, 0)
	for _, v := range list {
		reg := &types.UserFile{}
		reg.Id = v.Id
		reg.Ext = v.Ext
		reg.Identity = v.Identity
		reg.Name = v.Name
		reg.Path = v.Path
		reg.RepositoryIdentity = v.RepositoryIdentity
		reg.Size = v.Size
		res = append(res, reg)
	}

	resp.List = res
	resp.Count = cnt
	return resp, nil
}
