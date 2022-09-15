// Code generated by goctl. DO NOT EDIT!
// Source: email.proto

package server

import (
	"context"

	"core/app/email/cmd/rpc/internal/logic"
	"core/app/email/cmd/rpc/internal/svc"
	"core/app/email/cmd/rpc/pb"
)

type EmailCenterServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedEmailCenterServer
}

func NewEmailCenterServer(svcCtx *svc.ServiceContext) *EmailCenterServer {
	return &EmailCenterServer{
		svcCtx: svcCtx,
	}
}

func (s *EmailCenterServer) SendCode(ctx context.Context, in *pb.SendCodeRequest) (*pb.SendCodeResponse, error) {
	l := logic.NewSendCodeLogic(ctx, s.svcCtx)
	return l.SendCode(in)
}