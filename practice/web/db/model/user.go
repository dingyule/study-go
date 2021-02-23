package model

import (
	"fmt"
	"study-go/practice/web/db/utils"
)

type User struct {
	ID       int
	Username string
	Password string
	Email    string
}

//添加user的方法一
func (user *User) AddUser() error {
	//1、写sql语句
	sqlstr := "insert into users(username,password,email) values(?,?,?)"

	//2、预编译
	inStmt, err := utils.Db.Prepare(sqlstr)
	if err != nil {
		fmt.Println("编译出现异常", err)
		return err
	}
	//3、执行
	_, err2 := inStmt.Exec("admin", "123456", "test@admin.com")
	if err2 != nil {
		fmt.Println("执行出现异常", err2)
		return err2
	}
	return nil
}

//添加user的方法二
func (user *User) AddUser2() error {
	//1、写sql语句
	sqlstr := "insert into users(username,password,email) values(?,?,?)"

	//2、执行
	_, err := utils.Db.Exec(sqlstr, "admin2", "1111", "test2@admin.com")
	if err != nil {
		fmt.Println("执行出现异常", err)
		return err
	}
	return nil
}

//根据用户的id从数据库中查询一条记录
func (user *User) GetUserById() (*User, error) {
	//写sql语句
	sqlstr := "select * from users where id = ?"

	//执行
	row := utils.Db.QueryRow(sqlstr, user.ID)
	//声明
	var id int
	var username string
	var password string
	var email string
	err := row.Scan(&id, &username, &password, &email)
	if err != nil {
		return nil, err
	}
	u := &User{
		ID:       id,
		Username: username,
		Password: password,
		Email:    email,
	}
	return u, nil

}

//获取数据库中所有的记录
func (user *User) GetUsers() ([]*User, error) {
	//写sql语句
	sqlstr := "select * from users"

	//执行
	rows, err := utils.Db.Query(sqlstr)
	if err != nil {
		return nil, err
	}

	//创建user切片
	var users []*User
	for rows.Next() {
		//声明
		var id int
		var username string
		var password string
		var email string
		err := rows.Scan(&id, &username, &password, &email)
		if err != nil {
			return nil, err
		}
		u := &User{
			ID:       id,
			Username: username,
			Password: password,
			Email:    email,
		}
		users = append(users, u)
	}

	return users, nil

}
