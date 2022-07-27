package logic

import (
	"context"
	"fmt"

	"core/core/internal/svc"
	"core/core/internal/types"
	"core/models"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetShareBasicDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetShareBasicDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetShareBasicDetailLogic {
	return &GetShareBasicDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetShareBasicDetailLogic) GetShareBasicDetail(req *types.ShareBasicDetailRequest) (resp *types.ShareBasicDetailResponse, err error) {
	// 因为数据一致性要求不高，所以就不开启事务了
	// 看看资源存不存在
	// 分享资源的点击次数 +1
	engine := models.Engine
	share := new(models.ShareBasic)
	has, err := engine.Where("identity = ?", req.Identity).Get(share)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, fmt.Errorf("identity = %s 's share file is not existed ", req.Identity)
	}
	sql := "UPDATE share_basic SET click_num = click_num + 1 WHERE identity = ?"
	_, err = engine.Exec(sql, req.Identity)
	if err != nil {
		return nil, err
	}

	data := new(models.RepositoryPool)
	_, err = engine.Where("identity = ?", share.RepositoryIdentity).Get(data)
	if err != nil {
		return nil, err
	}
	resp = new(types.ShareBasicDetailResponse)
	resp.Name = data.Name
	resp.Ext = data.Ext
	resp.Path = data.Path
	resp.Size = data.Size
	resp.RepositoryIdentity = data.Identity
	return
}
