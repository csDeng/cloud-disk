package handler

import (
	"cloud_disk/app/internal/svc"
	"cloud_disk/app/response"
	"cloud_disk/logic"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func FileChunkSuccessHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FileChunkSuccessRequest
		if err := httpx.Parse(r, &req); err != nil {
			response.Response(w, nil, err)
			return
		}

		l := logic.NewFileChunkSuccessLogic(r.Context(), svcCtx)
		resp, err := l.FileChunkSuccess(&req)
		if err != nil {
			response.Response(w, nil, err)
		} else {
			response.Response(w, resp, err)
		}
	}
}
