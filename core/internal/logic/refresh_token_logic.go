package logic

import (
	"context"
	"time"

	"core/core/helper"
	"core/core/internal/svc"
	"core/core/internal/types"
	"core/redis"

	"github.com/zeromicro/go-zero/core/logx"
)

type RefreshTokenLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRefreshTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RefreshTokenLogic {
	return &RefreshTokenLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RefreshTokenLogic) RefreshToken(token string) (resp *types.RefreshTokenResponse, err error) {
	tokenCfg := helper.TokenConfigObject
	uc, err := helper.ParseToken(token)
	if err != nil {
		return nil, err
	}
	// 2. 生成 token
	token, err = helper.GenerateToken(uc.Id, uc.Identity, uc.Name)
	if err != nil {
		return nil, err
	}
	// 3. 生成refresh_token
	refresh_token, err := helper.GenerateToken(uc.Id, uc.Identity, uc.Name)
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
	resp = new(types.RefreshTokenResponse)
	resp.Token = token
	resp.RefreshToken = refresh_token
	return
}
