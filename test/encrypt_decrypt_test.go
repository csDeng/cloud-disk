package test

import (
	"core/core/helper"
	"testing"
)

func TestAesEncrypt(t *testing.T) {
	p := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJZCI6NiwiSWRlbnRpdHkiOiJjMjc4ZWRjNi0wYjJkLTQ3NzMtOWQ0Yy0xMzIzZDFhMjZjMjQiLCJOYW1lIjoibXluYW1lIiwiUmVmcmVzaFRva2VuSWQiOiJkZWNiNWE3ZS1kY2M1LTRhZGItYjkwNi1kY2FkMGNiOGEyODQiLCJleHAiOjE2NjA4OTQwNDd9.AAwMMGq48OuWsKSFz1lgWR9rJKAF7FBPIcW_xoXLy4M"
	t.Logf("%T %+v\r\n", p, p)
	c, err := helper.AesEncrypt(p)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("cipher=> %T %+v\r\n", c, c)
	pp, err := helper.AesDecrypt(c)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("plain=> %+v", pp)
	if p != pp {
		t.Fail()
	}
}
