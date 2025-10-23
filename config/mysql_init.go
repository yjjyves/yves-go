package config

import (
	"strings"

	//当导入一个包时，该包下的文件里所有init()函数都会被执行，然而，有些时候我们并不需要把整个包都导入进来，
	//仅仅是是希望它执行init()函数而已。这个时候就可以使用 import _ 引用该包
	/*
		使用了 . 作为前缀，表示在接下来调用该包中的变量、函数、类属性和成员方法时，无需使用包名前缀 animal. 引用，以免和 main 函数中的 animal 变量名冲突。
	*/
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var Db *sqlx.DB
var userName = "root"
var password = "root"
var address = "127.0.0.1:3306"
var databaseName = "test"

func InitDB(address string) (err error) {
	//dsn := "root:123456@tcp(127.0.0.1:3306)/go_db?charset=utf8mb4&parseTime=True"
	// 不会校验账号密码是否正确
	dataSourceName := getDataSourceName()
	database, err := sqlx.Open("mysql", dataSourceName)
	if err != nil {
		return err
	}
	// 尝试与数据库建立连接（校验dsn是否正确）
	err = database.Ping()
	if err != nil {
		return err
	}
	Db = database
	return nil
}

func getDataSourceName() string {
	var sb strings.Builder
	sb.WriteString(userName)
	sb.WriteString(":")
	sb.WriteString(password)
	sb.WriteString("@tcp")
	sb.WriteString(address)
	sb.WriteString(databaseName)

	return sb.String()
}
