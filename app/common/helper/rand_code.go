package helper

import (
	"math/rand"
	"time"
)

func RandCode(le int) string {
	s := "1234567890"
	code := ""
	rand.Seed(time.Now().Unix())
	for i := 0; i < le; i++ {
		code += string(s[rand.Intn(len(s))])
	}
	return code
}
