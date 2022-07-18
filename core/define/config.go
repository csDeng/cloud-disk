package define

import (
	"log"
	"os"
	"path/filepath"

	"github.com/go-ini/ini"
)

type EmailConfig struct {
	ServerName string `ini:"server"`
	Port       int    `ini:"port"`
	From       string `ini:"from"`
	Password   string `ini:"password"`
}

var EmailConfigObject = getConfig()

func getConfig() EmailConfig {
	str, err := os.Getwd()
	if err != nil {
		log.Fatal("get pwd fail, error: ", err)
	}
	str = filepath.Join(str, "../config/app.ini")
	cfg, err := ini.Load(str)
	if err != nil {
		log.Fatal("Fail to load config file, error: ", err)
	}
	config := EmailConfig{}
	err = cfg.Section("email").MapTo(&config)
	if err != nil {
		log.Fatal("fail to parse email config, error: ", err)
	}
	return config

}
