package test

import (
	"core/core/define"
	"core/core/helper"
	"fmt"
	"testing"
)

var EmailConfigObject define.EmailConfig

func TestEmailConfig(t *testing.T) {
	emailconfig := helper.EmailConfigObject
	fmt.Println(emailconfig)
}

func TestRedisConfig(t *testing.T) {
	rds := helper.RedisConfigObject
	fmt.Println(rds)
}
