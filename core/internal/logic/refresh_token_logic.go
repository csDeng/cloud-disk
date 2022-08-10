package logic

import (
	"context"
	"errors"
	"strconv"
	"time"

	"core/core/helper"
	"core/core/internal/svc"
	"core/core/internal/types"
	"core/redis"

	"github.com/zeromicro/go-zero/core/logx"

	grds "github.com/go-redis/redis/v8"
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
	rtk := helper.GetRefreshTokenKey(token)
	rds := redis.Redis
	// 检查当前的refresh_token 是否使用过
	val, err := rds.Get(l.ctx, rtk).Result()
	if err == grds.Nil {
		// do nothing
	} else if err != nil {
		return nil, err
	} else {
		v, err := strconv.Atoi(val)
		if err != nil {
			return nil, err
		}
		if v == 1 {
			return nil, errors.New("refresh_token 已被使用, 请确认refresh_token 是否为最新")
		}
	}

	uc, err := helper.ParseToken(token)
	if err != nil {
		return nil, err
	}
	// 2. 生成 token
	token, err = helper.GenerateToken(uc.Id, uc.Identity, uc.Name, false)
	if err != nil {
		return nil, err
	}
	// 3. 生成refresh_token
	refresh_token, err := helper.GenerateToken(uc.Id, uc.Identity, uc.Name, true)
	if err != nil {
		return nil, err
	}
	// 4. 标记refresh_token 已用过
	err = rds.SetEX(l.ctx, rtk, 1, time.Minute*time.Duration(tokenCfg.RefreshTokenTime)).Err()
	if err != nil {
		return nil, err
	}
	resp = new(types.RefreshTokenResponse)
	resp.Token = token
	resp.RefreshToken = refresh_token
	return
}
