package handler

import (
	"net/http"

	"core/core/define"
	"core/core/internal/logic"
	"core/core/internal/svc"
	"core/core/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func UserFileDelHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserFileDelRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewUserFileDelLogic(r.Context(), svcCtx)
		resp, err := l.UserFileDel(&req, r.Header.Get(define.HKOBJ.UserIdentity))
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
