package handler

import (
	"net/http"

	"cloud_disk/app/file/cmd/api/internal/logic"
	"cloud_disk/app/file/cmd/api/internal/svc"
	"cloud_disk/app/file/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func FileChunkUploadHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FileChunkUploadRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewFileChunkUploadLogic(r.Context(), svcCtx)
		resp, err := l.FileChunkUpload(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
