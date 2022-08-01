package handler

import (
	"errors"
	"net/http"
	"strconv"

	"core/core/helper"
	"core/core/internal/logic"
	"core/core/internal/svc"
	"core/core/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func FileChunkUploadHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FileChunkUploadRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}
		if len(r.PostForm.Get("key")) == 0 {
			err := errors.New("key 的长度不能为零")
			httpx.Error(w, err)
			return
		}
		if len(r.PostForm.Get("upload_id")) == 0 {
			err := errors.New("upload_id 的长度不能为零")
			httpx.Error(w, err)
			return
		}
		part_number, err := strconv.Atoi(r.PostForm.Get("part_number"))
		if err != nil {
			httpx.Error(w, err)
			return
		}
		if part_number == 0 {
			err := errors.New("part_number 不能为零")
			httpx.Error(w, err)
			return
		}
		eTag, err := helper.ChunkUpload(r)
		if err != nil {
			httpx.Error(w, err)
			return
		}
		l := logic.NewFileChunkUploadLogic(r.Context(), svcCtx)
		resp, err := l.FileChunkUpload(&req, eTag)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}