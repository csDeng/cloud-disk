package handler

import (
	"net/http"

	"cloud_disk/app/file/cmd/api/internal/logic"
	"cloud_disk/app/file/cmd/api/internal/svc"
	"cloud_disk/app/file/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func FileChunkSuccessHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FileChunkSuccessRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewFileChunkSuccessLogic(r.Context(), svcCtx)
		resp, err := l.FileChunkSuccess(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
