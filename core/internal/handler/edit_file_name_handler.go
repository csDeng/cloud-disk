package handler

import (
	"net/http"

	"core/core/define"
	"core/core/internal/logic"
	"core/core/internal/svc"
	"core/core/internal/types"
	"core/core/response"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func EditFileNameHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.EditFileNameRequest
		if err := httpx.Parse(r, &req); err != nil {
			response.Response(w, nil, err)
			return
		}

		l := logic.NewEditFileNameLogic(r.Context(), svcCtx)
		resp, err := l.EditFileName(&req, r.Header.Get(define.HKOBJ.UserIdentity))
		if err != nil {
			response.Response(w, nil, err)
		} else {
			response.Response(w, resp, err)
		}
	}
}
