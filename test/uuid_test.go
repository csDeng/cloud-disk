package test

import (
	"testing"

	"github.com/google/uuid"
)

func TestUuid(t *testing.T) {
	for i := 0; i < 10; i++ {
		uid := uuid.New().String()
		t.Logf("%T %v ", uid, uid)
		t.Log(len(uid))
	}
	t.Log(1)

}
