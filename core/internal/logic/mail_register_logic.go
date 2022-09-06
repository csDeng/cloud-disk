package logic

import (
	"context"
	"core/core/helper"
	"core/core/internal/svc"
	"core/core/internal/types"
	"core/models"
	"core/redis"
	"errors"
	"fmt"
	"time"

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
	// 判断邮箱是否存在
	cnt, err := models.Engine.Where("email = ?", req.Email).Count(new(models.UserBasic))
	if err != nil {
		return nil, err
	}
	// 如果已存在
	if cnt > 0 {
		return nil, errors.New("邮箱已存在")
	}
	rds := redis.Redis
	key := helper.GetMailRegKey(req.Email)

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
	code := helper.RandCode()

	err = helper.SendEmailCode(req.Email, code)
	if err != nil {
		return nil, err
	}
	_, err = rds.Set(l.ctx, key, code, time.Duration(helper.EmailConfigObject.Second)*time.Second).Result()
	if err != nil {
		return nil, err
	}
	resp = new(types.MailRegisterResponse)
	return resp, nil
}
