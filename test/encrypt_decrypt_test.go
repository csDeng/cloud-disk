package test

// import (
// 	"cloud_disk/app/helper"
// 	"testing"
// )

// func TestAesEncrypt(t *testing.T) {

// 	p := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJDaXBoZXIiOiJcdWZmZmRtRjNcdWZmZmReeFx1ZmZmZFx1ZmZmZFxuZVx1ZmZmZHBLXHVmZmZkQn9cdWZmZmRkXHVmZmZkXHUwMDNjdlx1MDAwZkJcdWZmZmTcpExcdTAwMTVcdWZmZmRqXHUwMDNlXHUwMDFkXHUwMDFjXHVmZmZkfEVcdWZmZmTXrj1cdWZmZmRsXHVmZmZkIiwiUmVmcmVzaFRva2VuSWQiOiI4NTRjNDdlOS0zMDZlLTQzOWQtOGQ2OS1hYmMxNWExZmQwNWUiLCJleHAiOjE2NjAxMzUyNTh9.EdKTA9whG4IIQ5ISrvV6oln8gteCeCmIwcCJRAB1P0s"
// 	t.Logf("%x", p)
// 	c, err := helper.AesEncrypt(p)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	t.Logf("cipher=> %+v", c)

// 	pp, err := helper.AesDecrypt(c)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	t.Logf("plain=> %+v", pp)
// 	if p != pp {
// 		t.Fail()
// 	}
// }
