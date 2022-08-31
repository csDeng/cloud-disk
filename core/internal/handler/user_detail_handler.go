package handler

import (
	"net/http"

	"core/core/internal/logic"
	"core/core/internal/svc"
	"core/core/internal/types"
	"core/core/response"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func UserDetailHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserDetailRequest
		if err := httpx.Parse(r, &req); err != nil {
			response.Response(w, nil, err)
			return
		}

		l := logic.NewUserDetailLogic(r.Context(), svcCtx)
		resp, err := l.UserDetail(&req)
		if err != nil {
			response.Response(w, nil, err)
		} else {
			response.Response(w, resp, err)
		}
	}
}
