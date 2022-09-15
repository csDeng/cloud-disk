package db

import (
	"core/app/common/vars"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

func GetEngine(mysql *vars.MysqlConfig) *xorm.Engine {
	host, port, user, password, dbname, charset := mysql.Server, mysql.Port, mysql.User, mysql.Password, mysql.DB, "utf8mb4"
	cnf := fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=%s",
		user, password, host, port, dbname, charset)
	engine, err := xorm.NewEngine("mysql", cnf)
	if err != nil {
		log.Printf("Xorm Engine error: %v \r\n", err)
		return nil
	}
	// 显示 sql 语句
	engine.ShowSQL(true)
	return engine
}
