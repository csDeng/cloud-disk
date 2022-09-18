package logic

import (
	"context"
	"log"
	"strconv"

	"core/app/common/helper"
	"core/app/common/vars"
	"core/app/user/cmd/rpc/internal/svc"
	"core/app/user/cmd/rpc/pb"
	"core/app/user/cmd/rpc/user_helper"

	"github.com/zeromicro/go-zero/core/logx"
)

type RefreshTokenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRefreshTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RefreshTokenLogic {
	return &RefreshTokenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RefreshTokenLogic) RefreshToken(in *pb.RefreshTokenRequest) (*pb.RefreshTokenResponse, error) {
	token := in.Token
	rds := l.svcCtx.RdsCli
	prefix := l.svcCtx.Config.RedisConf.Prefix
	rtk := helper.GetRefreshTokenKey(prefix, token)

	val, err := rds.Get(l.ctx, rtk).Result()
	log.Printf("get redis token, val=%v, \r\n err = %v \r\n ", val, err)
	if err == vars.ErrKeyIsNotExisted {
		return nil, vars.ErrRefreshTokenIsNotExisted
	} else if err != nil {
		return nil, err
	} else {
		v, err := strconv.Atoi(val)
		if err != nil {
			return nil, err
		}
		// 检查当前的refresh_token 是否在储存池
		if v == vars.TOKEN_STORE {
			// do nothing
		} else if v == vars.TOKEN_USED {
			return nil, vars.ErrRefreshTokenHasUsed
		}
	}

	uc, err := user_helper.ParseToken(token)
	if err != nil {
		return nil, err
	}
	// 2. 生成 token
	token, err = user_helper.GenerateToken(uc.Id, uc.Identity, uc.Name, false)
	if err != nil {
		return nil, err
	}
	// 3. 生成refresh_token
	refresh_token, err := user_helper.GenerateToken(uc.Id, uc.Identity, uc.Name, true)
	if err != nil {
		return nil, err
	}

	// 4. 将refresh_token 加入存储池
	// 5. 标记refresh_token 已用过
	new_rtk := helper.GetRefreshTokenKey(prefix, refresh_token)
	sha, err := user_helper.GetSha(l.ctx, rds)
	log.Println("get sha", sha, err)
	if err != nil {
		return nil, err
	}
	log.Printf("\r\n sha= %v \r\n", sha)
	v, err := rds.EvalSha(l.ctx, sha, []string{rtk, new_rtk, "Expire"},
		vars.TOKEN_USED, vars.TOKEN_STORE,
		60*l.svcCtx.Config.TokenConf.Refresh_token_time).Result()

	log.Printf("refresh_token=> %v \r\n err= %v", v, err)
	if err != nil || err == vars.ErrKeyIsNotExisted {
		return nil, err
	}
	if v.(int64) == vars.LUA_FAIL {
		return nil, vars.ErrLuaFail
	}

	// token 处理成功
	return &pb.RefreshTokenResponse{
		Token:        token,
		RefreshToken: refresh_token,
	}, nil
}
