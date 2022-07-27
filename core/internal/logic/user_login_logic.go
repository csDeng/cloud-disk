package logic

import (
	"context"
	"errors"
	"time"

	"core/core/helper"
	"core/core/internal/svc"
	"core/core/internal/types"
	"core/models"
	"core/redis"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLoginLogic {
	return &UserLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserLoginLogic) UserLogin(req *types.LoginRequest) (resp *types.LoginResponse, err error) {
	// 登录接口
	tokenCfg := helper.TokenConfigObject
	// 1. 从数据库获取登录用户信息

	engine := models.Engine
	user := new(models.UserBasic)
	has, err := engine.Where("name = ? AND password = ?", req.Name, helper.Md5(req.Password)).Get(user)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, errors.New("用户名或密码错误")
	}

	// 2. 生成 refresh_token
	refresh_token, err := helper.GenerateToken(user.Id, user.Identity, user.Name)
	if err != nil {
		return nil, err
	}
	// 3. 生成token
	token, err := helper.GenerateToken(user.Id, user.Identity, user.Name)
	if err != nil {
		return nil, err
	}
	// 4. 保存到redis
	tk := helper.GetTokenKey(token)
	rtk := helper.GetRefreshTokenKey(refresh_token)
	rds := redis.Redis

	err = rds.SetEX(l.ctx, tk, 1, time.Second*time.Duration(tokenCfg.TokenTime)).Err()
	if err != nil {
		return nil, err
	}
	err = rds.SetEX(l.ctx, rtk, 1, time.Second*time.Duration(tokenCfg.RefreshTokenTime)).Err()
	if err != nil {
		return nil, err
	}
	resp = new(types.LoginResponse)
	resp.Token = token
	resp.RefreshToken = refresh_token
	return

}
