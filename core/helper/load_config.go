package helper

import (
	"core/core/define"
	"log"
	"os"
	"path/filepath"

	"github.com/go-ini/ini"
)

// 获取 ini 文件
func GetConfigFile() *ini.File {
	pwd, err := os.Getwd()
	if err != nil {
		log.Fatal("get pwd fail, error: ", err)
	}
	path := filepath.Join(pwd, "../config/app.ini")
	if err != nil {
		log.Fatal("Fail to parse abs path, err: ", err)
	}
	cfg, err := ini.Load(path)
	if err != nil {
		log.Fatal("Fail to load config file, error: ", err)
	}
	return cfg
}

var MySqlConfigObject = getMysqlConfig()

func getMysqlConfig() *define.MysqlConfig {
	cfg := GetConfigFile()
	config := define.MysqlConfig{}
	err := cfg.Section("mysql").MapTo(&config)
	if err != nil {
		log.Fatal("fail to parse mysql config, error: ", err)
	}
	return &config
}

var EmailConfigObject = getEmailConfig()

func getEmailConfig() *define.EmailConfig {
	cfg := GetConfigFile()
	config := define.EmailConfig{}
	err := cfg.Section("email").MapTo(&config)
	if err != nil {
		log.Fatal("fail to parse email config, error: ", err)
	}
	return &config

}

var RedisConfigObject = getRedisConfig()

func getRedisConfig() *define.RedisConfig {
	cfg := GetConfigFile()

	config := define.RedisConfig{}
	err := cfg.Section("redis").MapTo(&config)
	if err != nil {
		log.Fatal("fail to parse redis config, error: ", err)
	}
	return &config
}

var CosConfigObject = getCosCfg()

func getCosCfg() *define.CosConfig {
	cfg := GetConfigFile()
	config := define.CosConfig{}
	err := cfg.Section("cos").MapTo(&config)
	if err != nil {
		log.Fatal("fail to parse cos config: error: ", err)
	}
	return &config
}
