package test

import (
	"core/core/helper"
	"fmt"
	"testing"
)

func TestMd5(t *testing.T) {
	s := "test12345"
	ss := helper.Md5(s)
	fmt.Println(s, ss)
}
