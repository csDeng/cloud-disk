package logic

import (
	"context"

	"core/core/helper"
	"core/core/internal/svc"
	"core/core/internal/types"

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


	
	err = helper.SendEmailCode(req.Email, "123456")
	if err != nil {
		return nil, err
	}
	return
}
