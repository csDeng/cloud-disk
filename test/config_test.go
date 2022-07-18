package test

import (
	"core/core/define"
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/go-ini/ini"
)

var EmailConfigObject define.EmailConfig

func TestEmailConfig(t *testing.T) {
	// 要传绝对路径
	str, err := os.Getwd()
	if err != nil {
		t.Fatal("get pwd fail, error: ", err)
	}
	str = filepath.Join(str, "../config/app.ini")
	cfg, err := ini.Load(str)
	if err != nil {
		t.Fatal("Fail to load config file, error: ", err)
	}
	EmailConfigObject = define.EmailConfig{}
	err = cfg.Section("email").MapTo(&EmailConfigObject)
	if err != nil {
		t.Fatal("fail to parse email config, error: ", err)
	}
	fmt.Println(EmailConfigObject)
}
