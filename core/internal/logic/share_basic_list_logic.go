package logic

import (
	"context"

	"core/core/internal/svc"
	"core/core/internal/types"
	"core/models"

	"github.com/zeromicro/go-zero/core/logx"
)

type ShareBasicListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewShareBasicListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShareBasicListLogic {
	return &ShareBasicListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ShareBasicListLogic) ShareBasicList(req *types.ShareBasicListRequest, userIdentity string) (resp *types.ShareBasicListResponse, err error) {
	engine := models.Engine
	share, repository := models.ShareBasicName, models.RepositoryPoolName
	page, size := req.Page+1, req.Size
	if page < 1 {
		page = 1
	}
	if size < 1 {
		size = 1
	}
	list := make([]*types.ShareFile, 0)
	offset := (page - 1) * size
	err = engine.Table(share).Alias("s").
		Select("r.name, r.ext, r.path, r.size,s.click_num").
		Join("left", []string{repository, "r"}, "s.repository_identity = r.Identity").
		Where("s.user_identity = ? AND s.deleted_at IS NULL", userIdentity).
		Where("r.deleted_at IS NULL").
		Limit(size, offset).Find(&list)
	if err != nil {
		return nil, err
	}

	cnt, err := engine.Where("user_identity = ? AND deleted_at IS NULL", userIdentity).Count(new(models.ShareBasic))
	if err != nil {
		return nil, err
	}
	resp = new(types.ShareBasicListResponse)
	resp.Count = int(cnt)
	resp.ShareBasic = list
	return
}
