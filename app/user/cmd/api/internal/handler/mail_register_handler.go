package handler

import (
	"net/http"

	"cloud_disk/app/user/cmd/api/internal/logic"
	"cloud_disk/app/user/cmd/api/internal/svc"
	"cloud_disk/app/user/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func MailRegisterHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.MailRegisterRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewMailRegisterLogic(r.Context(), svcCtx)
		resp, err := l.MailRegister(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
