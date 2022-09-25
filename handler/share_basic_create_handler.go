package handler

import (
	"cloud_disk/app/define"
	"cloud_disk/app/internal/logic"
	"cloud_disk/app/internal/svc"
	"cloud_disk/app/internal/types"
	"cloud_disk/app/response"
	"errors"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func ShareBasicCreateHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ShareBasicCreateRequest
		if err := httpx.Parse(r, &req); err != nil {
			response.Response(w, nil, err)
			return
		}

		if len(req.RepositoryIdentity) < 1 {
			httpx.Error(w, errors.New("repository_identity 长度不能小于1"))
			return
		}

		l := logic.NewShareBasicCreateLogic(r.Context(), svcCtx)
		resp, err := l.ShareBasicCreate(&req, r.Header.Get(define.HKOBJ.UserIdentity))
		if err != nil {
			response.Response(w, nil, err)
		} else {
			response.Response(w, resp, err)
		}
	}
}
