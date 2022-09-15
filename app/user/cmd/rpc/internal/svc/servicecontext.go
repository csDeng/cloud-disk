package svc

import (
	"core/app/common/db"
	"core/app/common/helper"
	"core/app/common/rds"
	"core/app/common/vars"
	"core/app/email/cmd/rpc/emailcenter"
	"core/app/user/cmd/rpc/internal/config"
	"core/app/user/model"
	"log"

	"github.com/go-redis/redis/v8"
	"github.com/zeromicro/go-zero/zrpc"
	"xorm.io/xorm"
)

var engine *xorm.Engine

func getEngine(c *config.Config) *xorm.Engine {
	if engine == nil {
		MySqlConfigObject := vars.MysqlConfig{
			Server:   c.Mysql.Server,
			Port:     c.Mysql.Port,
			User:     c.Mysql.User,
			Password: c.Mysql.Password,
			DB:       c.Mysql.Db,
		}
		engine = db.GetEngine(&MySqlConfigObject)
	}
	return engine
}

var rdscli *redis.Client

func getRdsCli(c *config.Config) *redis.Client {
	if rdscli == nil {
		RedisConfigObject := vars.RedisConfig{
			Server:      c.RedisConf.Server,
			Port:        c.RedisConf.Port,
			RedisPrefix: c.RedisConf.Prefix,
		}
		rdscli = rds.GetRdsClient(&RedisConfigObject)
	}
	return rdscli
}

type ServiceContext struct {
	Config config.Config

	EmailRpcClient emailcenter.EmailCenter

	UserModel *model.UserBasic

	Engine *xorm.Engine

	RdsCli *redis.Client

	// 验证码长度
	RandCodeLength int
}

func NewServiceContext(c config.Config) *ServiceContext {
	// 配置注入
	aesConf := &vars.AesCfg{
		Secret: c.AesConf.Secret,
		IV:     c.AesConf.Iv,
	}

	tokenConf := &vars.TokenConfig{
		TokenTime:        c.TokenConf.Token_time,
		RefreshTokenTime: c.TokenConf.Refresh_token_time,
		Secret:           c.TokenConf.Secret,
	}

	helper.InjectAesCfg(aesConf)
	helper.InjectTokenCfg(tokenConf)

	log.Println("注入配置")

	return &ServiceContext{
		Config:         c,
		EmailRpcClient: emailcenter.NewEmailCenter(zrpc.MustNewClient(c.EmailRpcConf)),
		UserModel:      model.NewUserModel(),
		Engine:         getEngine(&c),

		RdsCli: getRdsCli(&c),

		RandCodeLength: c.RandCodeLength,
	}
}
