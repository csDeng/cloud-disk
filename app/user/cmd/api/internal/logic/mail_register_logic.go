package logic

import (
	"context"
	"fmt"

	"core/app/user/cmd/api/internal/svc"
	"core/app/user/cmd/api/internal/types"
	"core/app/user/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type MailRegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMailRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MailRegisterLogic {
	return &MailRegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MailRegisterLogic) MailRegister(req *types.MailRegisterRequest) (resp *types.MailRegisterResponse, err error) {
	// todo: add your logic here and delete this line
	res, err := l.svcCtx.UserRpcClient.Register(l.ctx, &pb.UserRegisterRequest{
		Code:     "123456",
		Email:    req.Email,
		Name:     "",
		Password: "",
	})
	fmt.Println("======================", req.Email)
	if err != nil {
		return nil, err
	}
	fmt.Println(res)
	return
}
