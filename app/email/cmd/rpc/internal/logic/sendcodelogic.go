package logic

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"net/smtp"

	"cloud_disk/app/email/cmd/rpc/internal/svc"
	"cloud_disk/app/email/cmd/rpc/pb"

	"github.com/jordan-wright/email"
	"github.com/zeromicro/go-zero/core/logx"
)

type SendCodeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSendCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendCodeLogic {
	return &SendCodeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SendCodeLogic) SendCode(in *pb.SendCodeRequest) (*pb.SendCodeResponse, error) {
	mail, code := in.Email, in.Code

	emailConfig := l.svcCtx.EmailConfigObject
	resp := new(pb.SendCodeResponse)

	e := email.NewEmail()
	e.From = fmt.Sprintf("<%s>", emailConfig.From)
	e.To = []string{mail}
	e.Subject = "邮箱注册"
	e.HTML = []byte("你的验证码是: " + code)
	err := e.SendWithTLS(
		fmt.Sprintf("%s:%d", emailConfig.ServerName, emailConfig.Port),
		smtp.PlainAuth("", emailConfig.From, emailConfig.Password, emailConfig.ServerName),
		&tls.Config{InsecureSkipVerify: true, ServerName: emailConfig.ServerName})

	if err != nil {
		log.Printf("邮箱注册: %s, 验证码发送失败, error = %v\r\n", mail, err)
		return resp, err
	}
	return resp, nil
}
