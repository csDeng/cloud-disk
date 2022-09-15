package user_helper

import (
	"core/app/common/vars"
	"log"
)

func InitCfg(aesCfg *vars.AesCfg, tokenCfg *vars.TokenConfig) {
	InitAesCfg(aesCfg)
	InitTokenCfg(tokenCfg)
	log.Println("配置注入", aesCfg, tokenCfg)
}
