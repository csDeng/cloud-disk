package helper

import (
	"core/core/define"
	"math/rand"
	"time"
)

func RandCode() string {
	s := "1234567890"
	le := define.Code_length
	code := ""
	rand.Seed(time.Now().Unix())
	for i := 0; i < le; i++ {
		code += string(s[rand.Intn(len(s))])
	}
	return code
}
