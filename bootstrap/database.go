package bootstrap

import (
	"errors/pkg/database"
	"fmt"
	"time"
)

func SetupDB() {

	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=%v&parseTime=True&multiStatements=true&loc=Local",
		"root",
		"root",
		"127.0.0.1",
		3306,
		"database",
		"utf8mb4",
	)
	database.Connection(dsn)
	// 设置最大连接数
	database.DB.SetMaxOpenConns(30)
	// 设置最大空闲连接数
	database.DB.SetMaxIdleConns(20)
	// 设置每个链接的过期时间
	database.DB.SetConnMaxLifetime(time.Duration(5*60) * time.Second)
}
