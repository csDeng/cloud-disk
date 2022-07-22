package logic

import (
	"context"
	"errors"
	"fmt"

	"core/core/helper"
	"core/core/internal/svc"
	"core/core/internal/types"
	"core/models"
	"core/rds"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserRegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRegisterLogic {
	return &UserRegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserRegisterLogic) UserRegister(req *types.UserRegisterRequest) (resp *types.UserRegisterResponse, err error) {
	key := helper.GetMailRegKey(req.Email)
	redis := rds.Redis
	v, err := redis.Get(l.ctx, key).Result()
	if err != nil {
		return nil, errors.New("当前邮箱没有获取验证码!")
	}
	if v != req.Code {
		return nil, errors.New("验证码错误!")
	}
	user := new(models.UserBasic)
	cnt, err := models.Engine.Where("name = ?", req.Name).Count(user)
	if err != nil {
		return nil, err
	}
	if cnt > 0 {
		return nil, errors.New(fmt.Sprintf("用户名: %s 已存在", req.Name))
	}
	user.Identity = helper.GenerateUuid()
	user.Name = req.Name
	user.Password = helper.Md5(req.Password)
	user.Email = req.Email
	u, err := models.Engine.InsertOne(user)
	if err != nil {
		return nil, err
	}
	if u > 0 {
		_, err = models.Engine.Where("identity = ?", user.Identity).Get(user)
		if err != nil {
			return nil, err
		}
		resp := new(types.UserRegisterResponse)
		resp.Identity = user.Identity
		resp.Email = user.Email
		resp.Name = user.Name
		return resp, nil
	}
	return nil, nil
}
