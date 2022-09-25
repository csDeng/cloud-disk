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

func FileUploadHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FileUploadRequest
		if err := httpx.Parse(r, &req); err != nil {
			response.Response(w, nil, err)
			return
		}
		id, err := svcCtx.UserRpcClient.GetIdentityWithToken(r.Context(),
			&pb.GetIdentityWithTokenRequest{
				Token: r.Header.Get(vars.Header_TOKEN),
			})
		if err != nil {
			response.Response(w, nil, err)
		}
		l := logic.NewFileUploadLogic(r.Context(), svcCtx)
		resp, err := l.FileUpload(&req, id.Identity, r)
		response.Response(w, resp, err)
	}

}
