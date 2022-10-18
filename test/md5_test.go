package test

import (
	"cloud_disk/app/common/helper"
	"fmt"
	"testing"
)

// import (
// 	"cloud_disk/app/helper"
// 	"crypto/md5"
// 	"fmt"
// 	"testing"
// )

func TestMd5(t *testing.T) {
	s := "test123456"
	ss := helper.Md5(s)
	fmt.Println(ss)
}

// func TestMd5Sum(t *testing.T) {
// 	secret := "hello"
// 	sum := md5.Sum([]byte(secret))
// 	s := sum[:]
// 	t.Logf("%v", sum)
// 	t.Logf("%#v", s)
// }
