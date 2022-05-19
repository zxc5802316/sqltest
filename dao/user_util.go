package dao

import (
	"errors/pkg/database"
	"github.com/pkg/errors"
)

// QueryUserByName 预处理 没有 sql.ErrNoRows
func QueryUserByName(name string) ([]User, error) {

	prepare, err := database.DB.Prepare("select `id`, `name`, `age` from user where `name` = ?")
	if err != nil {
		return nil, errors.Wrap(err, "dao: 预处理失败")
	}
	defer prepare.Close()
	rows, err := prepare.Query(name)
	if err != nil {
		return nil, errors.Wrap(err, "dao: sql 执行失败")
	}
	defer rows.Close()
	var users []User
	for rows.Next() {
		user := User{}
		err := rows.Scan(&user.ID, &user.Name, &user.Age)
		if err != nil {
			return nil, errors.Wrap(err, "dao: 查询失败")
		}
		users = append(users, user)
	}
	return users, nil
}

func QueryUserByName1(name string) (*User, error) {

	sql := "select `id`, `name`, `age` from user where `name` = ?"
	user := User{}
	rows := database.DB.QueryRow(sql, name)
	err := rows.Scan(&user.ID, &user.Name, &user.Age)
	if err != nil {
		return nil, errors.Wrapf(err, "dao: QueryUserByName1 sql := %s name = %s \n", sql, name)
	}
	return &user, nil

}
func Insert(user User) error {
	t, err := database.DB.Begin()
	if err != nil {
		return errors.Wrap(err, "dao: 事务失败")
	}
	prepare, err := t.Prepare("insert into user (`name`,`age`) values (?,?)")
	if err != nil {
		return errors.Wrap(err, "dao: 预处理失败")
	}
	_, err = prepare.Exec(user.Name, user.Age)
	if err != nil {
		return errors.Wrap(err, "dao: sql 执行失败")
	}
	err = t.Commit()
	if err != nil {
		return errors.Wrap(err, "dao: 数据库事务Commit失败")
	}
	return nil
}
