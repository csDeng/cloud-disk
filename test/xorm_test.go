package test

import (
	"bytes"
	"core/models"
	"encoding/json"
	"fmt"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

func TestXormTest(t *testing.T) {
	user, password, dbname, charset := "root", "root", "cloud_disk", "utf8mb4"
	cnf := fmt.Sprintf("%s:%s@/%s?charset=%s",
		user, password, dbname, charset)
	engine, err := xorm.NewEngine("mysql", cnf)
	if err != nil {
		t.Fatal(err)
	}

	data := make([]*models.UserBasic, 0)
	err = engine.Find(&data)
	if err != nil {
		t.Fatal(err)
	}

	// 转换成json
	b, err := json.Marshal(data)
	if err != nil {
		t.Fatal(err)
	}
	dst := new(bytes.Buffer)

	// 添加缩进
	err = json.Indent(dst, b, "", "    ")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(dst.String())

}

func TestInsert(t *testing.T) {
	user, password, dbname, charset := "root", "root", "cloud_disk", "utf8mb4"
	cnf := fmt.Sprintf("%s:%s@/%s?charset=%s",
		user, password, dbname, charset)
	engine, err := xorm.NewEngine("mysql", cnf)
	if err != nil {
		t.Fatal(err)
	}

	v, err := engine.Insert(&models.UserBasic{
		Name:     "username",
		Password: "pwd",
		Email:    "99@99.com",
	})
	if err != nil {
		t.Fatal(err)
	}
	res := new(models.UserBasic)
	if v > 0 {
		engine.Where("name = ?", "username").And("email = ?", "99@99.com").Get(res)
	}
	t.Log(res)
}
