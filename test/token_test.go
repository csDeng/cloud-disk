package test

import (
	"core/core/helper"
	"testing"
)

func TestToken(t *testing.T) {
	token, err := helper.GenerateToken(1, "identity", "name", false)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%v, %T", token, token)
	uc, err := helper.ParseToken(token)
	if err != nil {
		t.Log(err)
	}
	t.Logf("%#v", uc)
}
