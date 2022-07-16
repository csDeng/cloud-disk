package models

import (
	"fmt"
	"log"

	// 调用 mysql 的 init 函数
	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

// 获取Engine 的时候调用 Init函数
var Engine = getEngine()

func getEngine() *xorm.Engine {
	user, password, dbname, charset := "root", "root", "cloud_disk", "utf8mb4"
	cnf := fmt.Sprintf("%s:%s@/%s?charset=%s",
		user, password, dbname, charset)
	engine, err := xorm.NewEngine("mysql", cnf)
	if err != nil {
		log.Printf("Xorm Engine error: %v \r\n", err)
		return nil
	}
	return engine
}
