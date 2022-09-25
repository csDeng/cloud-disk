package logic

import (
	"context"
	"fmt"
	"time"

	"cloud_disk/app/common/helper"
	"cloud_disk/app/common/vars"
	"cloud_disk/app/email/cmd/rpc/emailcenter"
	"cloud_disk/app/user/cmd/rpc/internal/svc"
	"cloud_disk/app/user/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type EmailVerificationLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewEmailVerificationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EmailVerificationLogic {
	return &EmailVerificationLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *EmailVerificationLogic) EmailVerification(in *pb.EmailVerificationRequest) (*pb.EmailVerificationResponse, error) {
	b, err := l.svcCtx.UserModel.EmailIfExisted(l.svcCtx.Engine, in.Email)
	if err != nil {
		return nil, err
	}
	// 邮箱已存在
	if b {
		return nil, vars.ErrEmailIsExisted
	}

	rds := l.svcCtx.RdsCli
	prefix := l.svcCtx.Config.RedisConf.Prefix
	key := helper.GetMailRegKey(prefix, in.Email)
	// 判断是否以获取过随机码
	v, err := rds.Exists(l.ctx, key).Result()
	if err != nil {
		return nil, err
	}
	if v > 0 {
		vv, err := rds.Do(l.ctx, "TTL", key).Result()
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("验证码已发送,请 %d s 后重试", vv.(int64))
	}

	// 获取随机码
	code := helper.RandCode(l.svcCtx.RandCodeLength)

	_, err = l.svcCtx.EmailRpcClient.SendCode(l.ctx, &emailcenter.SendCodeRequest{
		Email: in.Email,
		Code:  code,
	})
	if err != nil {
		return nil, vars.ErrEmailSend
	}

	err = rds.Set(l.ctx, key, code, time.Duration((l.svcCtx.Config.RandCodeExpire))*time.Second).Err()

	if err != nil {
		return nil, vars.ErrRegCode
	}

	return &pb.EmailVerificationResponse{}, nil
}
