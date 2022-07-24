package handler

import (
	"crypto/md5"
	"fmt"
	"net/http"
	"path"

	"core/core/helper"
	"core/core/internal/logic"
	"core/core/internal/svc"
	"core/core/internal/types"
	"core/models"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func FileUploadHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FileUploadRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}
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

		rp := new(models.RepositoryPool)
		has, err := models.Engine.Where("hash = ?", hash).Get(rp)
		if err != nil {
			return
		}
		if has {
			httpx.OkJson(w, &types.FileUploadResponse{
				Identity: rp.Identity,
				Name:     rp.Name,
				Ext:      rp.Ext,
				Size:     int64(rp.Size),
				Path:     rp.Path,
			})
			return
		}

		// 如果文件不存在，才上传
		pathString, err := helper.UploadFile(r)
		if err != nil {
			return
		}

		req.Name = fileHeader.Filename
		req.Ext = path.Ext(fileHeader.Filename)
		req.Size = fileHeader.Size
		req.Hash = hash
		req.Path = pathString

		l := logic.NewFileUploadLogic(r.Context(), svcCtx)
		resp, err := l.FileUpload(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
