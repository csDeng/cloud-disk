package logic

import (
	"context"
	"errors"

	"cloud_disk/app/common/helper"
	"cloud_disk/app/common/vars"
	"cloud_disk/app/user/cmd/rpc/internal/svc"
	"cloud_disk/app/user/cmd/rpc/pb"
	"cloud_disk/app/user/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *pb.UserRegisterRequest) (*pb.UserRegisterResponse, error) {
	userModel, Engine := l.svcCtx.UserModel, l.svcCtx.Engine
	b, err := userModel.EmailIfExisted(Engine, in.Email)
	if err != nil {
		return nil, err
	}
	if b {
		return nil, vars.ErrEmailIsExisted
	}
	rds := l.svcCtx.RdsCli
	prefix := l.svcCtx.Config.RedisConf.Prefix
	key := helper.GetMailRegKey(prefix, in.Email)

	v, err := rds.Get(l.ctx, key).Result()
	if err != nil {
		return nil, err
	} else if err == vars.ErrKeyIsNotExisted {
		return nil, vars.ErrEmailNotGetCode
	}
	if v != in.Code {
		return nil, errors.New("验证码错误")
	}
	identity := helper.GenerateUuid()
	newUser := &model.UserBasic{
		Identity: identity,
		Name:     in.Name,
		Password: helper.Md5(in.Password),
		Email:    in.Email,
	}

	b, err = userModel.AddUser(Engine, newUser)
	if err != nil {
		return nil, err
	}
	if !b {
		return nil, vars.ErrAdd
	}

	user, err := userModel.GetUserByIdentity(Engine, identity)

	if err != nil {
		return nil, err
	}

	resp := new(pb.UserRegisterResponse)
	resp.Name = user.Name
	resp.Identity = user.Identity
	resp.Email = user.Email

	return resp, nil
}
