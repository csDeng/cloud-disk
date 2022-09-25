package svc

import (
	"cloud_disk/app/common/db"
	"cloud_disk/app/common/vars"
	"cloud_disk/app/file/cmd/api/internal/config"
	"cloud_disk/app/file/models"
	"cloud_disk/app/user/cmd/rpc/usercenter"

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

type ServiceContext struct {
	Config config.Config

	UserRpcClient usercenter.UserCenter

	Engine *xorm.Engine

	PoolModel      *models.RepositoryPool
	UserRepository *models.UserRepository
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		UserRpcClient: usercenter.NewUserCenter(
			zrpc.MustNewClient(c.UserRpcConf),
		),

		Engine: getEngine(&c),

		PoolModel: models.NewPoolModel(),

		UserRepository: models.NewUserRepository(),
	}
}
