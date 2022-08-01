package handler

import (
	"errors"
	"net/http"

	"core/core/internal/logic"
	"core/core/internal/svc"
	"core/core/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func FileUploadPrepareHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FileUploadPrepareRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		if len(req.Md5) == 0 {
			httpx.Error(w, errors.New("md5 长度不能为零"))
			return
		}
		if len(req.Name) == 0 {
			httpx.Error(w, errors.New("name 长度不能为零"))
			return
		}
		if len(req.Ext) == 0 {
			httpx.Error(w, errors.New("ext 长度不能为零"))
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
