package svc

import (
	"cloud_disk/app/common/db"
	"cloud_disk/app/common/rds"
	"cloud_disk/app/common/vars"
	"cloud_disk/app/email/cmd/rpc/emailcenter"
	"cloud_disk/app/user/cmd/rpc/internal/config"
	"cloud_disk/app/user/model"

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
			Password:    c.RedisConf.Password,
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
	return &ServiceContext{
		Config:         c,
		EmailRpcClient: emailcenter.NewEmailCenter(zrpc.MustNewClient(c.EmailRpcConf)),
		UserModel:      model.NewUserModel(),
		Engine:         getEngine(&c),

		RdsCli: getRdsCli(&c),

		RandCodeLength: c.RandCodeLength,
	}
}
