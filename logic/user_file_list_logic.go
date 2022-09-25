package logic

import (
	"context"

	"cloud_disk/app/internal/svc"
	"cloud_disk/app/internal/types"

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
	// resp = new(types.UserFileListResponse)
	// uf := make([]*types.UserFile, 0)
	// size := req.Size
	// if size == 0 {
	// 	size = define.PageSize
	// }
	// page := req.Page
	// if page <= 0 {
	// 	page = 1
	// }
	// offset := (page - 1) * size
	/**
	type UserRepository struct {
		Id                 int    `xorm:"id"`
		Identity           string `xorm:"identity"`
		ParentId           int    `xorm:"parent_id"`
		UserIdentity       string `xorm:"user_identity"`
		RepositoryIdentity string `xorm:"repository_identity"`
		Ext                string `xorm:"ext"`
		Name               string `xorm:"name"`


	}

	type RepositoryPool struct {
		Id       int    `xorm:"'id' pk"`
		Identity string `xorm:"'identity'"`
		Size     int    `xorm:"'size'"`
		Path     string `xorm:"'path'"`

	}
	type UserFileListResponse struct {
		List  []*UserFile `json:"list"`
		Count int         `json:"count"`
	}

	type UserFile struct {
		Id                 int    `json:"id"`
		Name               string `json:"name"`
		Ext                string `json:"ext"`
		Size               int    `json:"size"`
		Path               string `json:"path"`
		Identity           string `json:"identity"`
		RepositoryIdentity string `json:"repository_identity"`
	}

			**/
	// urName, pName := models.UserRepositoryName, models.RepositoryPoolName
	// err = models.Engine.Table(urName).
	// 	Where("parent_id = ? AND user_identity = ?", req.Id, userIdentity).
	// 	Select("user_repository.id, user_repository.name, user_repository.ext, repository_pool.size,repository_pool.path,user_repository.identity, repository_pool.identity AS repository_identity").
	// 	Join("LEFT", pName, "user_repository.repository_identity = repository_pool.Identity").
	// 	Where("user_repository.deleted_at IS NULL").
	// 	Limit(size, offset).Find(&uf)
	// if err != nil {
	// 	return nil, err
	// }
	// cnt, err := models.Engine.Where("parent_id = ? AND user_identity = ?", req.Id, userIdentity).
	// 	Count(new(models.UserRepository))
	// if err != nil {
	// 	return nil, err
	// }
	// resp.List = uf
	// resp.Count = int(cnt)

	return
}
