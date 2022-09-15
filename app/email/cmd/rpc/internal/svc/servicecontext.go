package svc

import (
	"core/app/common/vars"
	"core/app/email/cmd/rpc/internal/config"
	"fmt"
)

type ServiceContext struct {
	Config            config.Config
	EmailConfigObject vars.EmailConfig
}

func NewServiceContext(c config.Config) *ServiceContext {
	fmt.Println(c.EmailConfig)
	return &ServiceContext{
		Config: c,
		EmailConfigObject: vars.EmailConfig{
			ServerName: c.EmailConfig.Server,
			Port:       c.EmailConfig.Port,
			From:       c.EmailConfig.From,
			Password:   c.EmailConfig.Password,
			Second:     c.EmailConfig.Second,
		},
	}
}
