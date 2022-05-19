package main

import (
	"database/sql"
	"errors/bootstrap"
	"errors/dao"
	"fmt"
	"github.com/pkg/errors"
)

func main() {
	bootstrap.SetupDB()
	//user := dao.User{
	//	Name: "zhang",
	//	Age :17,
	//}
	//err := dao.Insert(user)
	//if err != nil {
	//	fmt.Printf("cause error : %+v \n",errors.Cause(err))
	//	fmt.Printf("stack : \n %+v \n",err)
	//}
	//
	users, err := dao.QueryUserByName1("ii")
	if err != nil {
		fmt.Printf("stack : \n %+v \n", err)
		if errors.Is(err, sql.ErrNoRows) {
			doSomething()
		}
		return
	}
	fmt.Printf("%v", users)
}

func doSomething() {
	//	记录日志
	//	以及处理其他操作
}
