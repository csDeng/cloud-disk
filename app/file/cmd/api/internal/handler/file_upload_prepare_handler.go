package handler

import (
	"net/http"

	"cloud_disk/app/file/cmd/api/internal/logic"
	"cloud_disk/app/file/cmd/api/internal/svc"
	"cloud_disk/app/file/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func FileUploadPrepareHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FileUploadPrepareRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewFileUploadPrepareLogic(r.Context(), svcCtx)
		resp, err := l.FileUploadPrepare(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
