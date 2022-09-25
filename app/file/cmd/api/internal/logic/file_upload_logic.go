package logic

import (
	"context"
	"crypto/md5"
	"fmt"
	"log"
	"net/http"
	"path"

	"cloud_disk/app/common/helper"
	"cloud_disk/app/file/cmd/api/internal/svc"
	"cloud_disk/app/file/cmd/api/internal/types"
	"cloud_disk/app/file/file_helper"
	"cloud_disk/app/file/models"

	"github.com/zeromicro/go-zero/core/logx"
)

type FileUploadLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFileUploadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FileUploadLogic {
	return &FileUploadLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FileUploadLogic) FileUpload(req *types.FileUploadRequest, user_identity string, r *http.Request) (resp *types.FileUploadResponse, err error) {
	file, fileHeader, err := r.FormFile("file")
	if err != nil {
		return
	}

	// 判断文件是否已存在
	b := make([]byte, fileHeader.Size)
	// 往 byte 里面写数据
	_, err = file.Read(b)
	if err != nil {
		return
	}
	hash := fmt.Sprintf("%x", md5.Sum(b))
	PoolModel, Engine := l.svcCtx.PoolModel, l.svcCtx.Engine
	rp, err := PoolModel.CheckFileIfExisted(Engine, hash)
	if err != nil {
		return
	}
	ext := path.Ext(fileHeader.Filename)
	if len(rp.Identity) == 0 {
		// 将文件存进池子
		p, err := file_helper.CreateFile(b, ext)
		if err != nil {
			return nil, err
		}

		rp = &models.RepositoryPool{
			Identity: helper.GenerateUuid(),
			Hash:     hash,
			Name:     fileHeader.Filename,
			Ext:      ext,
			Size:     int(fileHeader.Size),
			Path:     p,
		}
		_, err = PoolModel.AddFile(Engine, rp)
		if err != nil {
			return nil, err
		}
	}
	log.Println("rp=>", rp)
	UserRepositoryModel := l.svcCtx.UserRepository
	resp = new(types.FileUploadResponse)

	ur, err := UserRepositoryModel.CheckUserIfHasFile(Engine, user_identity, rp.Identity)
	if err != nil {
		return nil, err
	}
	// 说明文件没有与用户关联存储
	if len(ur.Identity) == 0 {
		ur, err = UserRepositoryModel.AddFile(Engine, &models.UserRepository{
			Identity:           helper.GenerateUuid(),
			Ext:                rp.Ext,
			Name:               rp.Name,
			UserIdentity:       user_identity,
			RepositoryIdentity: rp.Identity,
		})
		if err != nil {
			return nil, err
		}
	}
	resp.Identity = ur.RepositoryIdentity
	resp.Ext = ur.Ext
	resp.Name = ur.Name
	resp.Size = int64(rp.Size)
	resp.Path = rp.Path
	return resp, nil
}
