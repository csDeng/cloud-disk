// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"core/core/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/user/login",
				Handler: UserLoginHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/user/detail",
				Handler: UserDetailHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/mail/register",
				Handler: MailRegisterHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/share",
				Handler: GetShareBasicDetailHandler(serverCtx),
			},
			{
				Method:  http.MethodPut,
				Path:    "/refresh",
				Handler: RefreshTokenHandler(serverCtx),
			},
		},
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.AuthMiddleware},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/user/register",
					Handler: UserRegisterHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/file/upload",
					Handler: FileUploadHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/user/repository/save",
					Handler: UserRepositorySaveHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/user/repository",
					Handler: UserFileListHandler(serverCtx),
				},
				{
					Method:  http.MethodPut,
					Path:    "/user/repository",
					Handler: EditFileNameHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/user/folder",
					Handler: UserFolderCreateHandler(serverCtx),
				},
				{
					Method:  http.MethodDelete,
					Path:    "/user/file",
					Handler: UserFileDelHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/user/file/move",
					Handler: UserFileMoveHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/share",
					Handler: ShareBasicCreateHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/share/list",
					Handler: ShareBasicListHandler(serverCtx),
				},
			}...,
		),
	)
}
