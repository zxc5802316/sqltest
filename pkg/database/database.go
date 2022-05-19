package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var DB  *sql.DB

func Connection(dsn string)  {
	var err error
	DB, err = sql.Open("mysql", dsn)
	if err!=nil {
		fmt.Printf("dao: 数据库参数有误 dns: %s err: %v\n", dsn, err)
		return
	}
	err = DB.Ping()
	if err != nil {
		fmt.Printf("dao: 数据库连接失败 dns: %s err: %v\n", dsn, err)
		return
	}
}