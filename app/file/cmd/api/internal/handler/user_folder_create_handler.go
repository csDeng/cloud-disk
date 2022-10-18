package handler

import (
	"net/http"

	"cloud_disk/app/common/response"
	"cloud_disk/app/common/vars"
	"cloud_disk/app/file/cmd/api/internal/logic"
	"cloud_disk/app/file/cmd/api/internal/svc"
	"cloud_disk/app/file/cmd/api/internal/types"
	"cloud_disk/app/user/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func UserFolderCreateHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserFolderCreateRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		id, err := svcCtx.UserRpcClient.GetIdentityWithToken(r.Context(),
			&pb.GetIdentityWithTokenRequest{
				Token: r.Header.Get(vars.Header_TOKEN),
			})
		if err != nil {
			response.Response(w, nil, err)
		}

		l := logic.NewUserFolderCreateLogic(r.Context(), svcCtx)

		resp, err := l.UserFolderCreate(&req, id.Identity)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
