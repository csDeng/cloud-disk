package handler

import (
	"net/http"

	"cloud_disk/app/define"
	"cloud_disk/app/internal/logic"
	"cloud_disk/app/internal/svc"
	"cloud_disk/app/internal/types"
	"cloud_disk/app/response"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func ShareBasicListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ShareBasicListRequest
		if err := httpx.Parse(r, &req); err != nil {
			response.Response(w, nil, err)
			return
		}

		l := logic.NewShareBasicListLogic(r.Context(), svcCtx)
		resp, err := l.ShareBasicList(&req, r.Header.Get(define.HKOBJ.UserIdentity))
		if err != nil {
			response.Response(w, nil, err)
		} else {
			response.Response(w, resp, err)
		}
	}
}
