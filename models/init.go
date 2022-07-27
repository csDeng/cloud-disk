package models

import (
	"core/core/helper"
	"fmt"
	"log"

	// 调用 mysql 的 init 函数
	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

// 获取Engine 的时候调用 Init函数
var Engine = getEngine()

func getEngine() *xorm.Engine {
	mysql := helper.MySqlConfigObject
	host, port, user, password, dbname, charset := mysql.Server, mysql.Port, mysql.User, mysql.Password, mysql.DB, "utf8mb4"
	cnf := fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=%s",
		user, password, host, port, dbname, charset)
	engine, err := xorm.NewEngine("mysql", cnf)
	// 显示 sql 语句
	engine.ShowSQL(true)
	if err != nil {
		log.Printf("Xorm Engine error: %v \r\n", err)
		return nil
	}
	return engine
}
