package model

import (
	"fmt"
	"testing"
)

//func TestAddUser(t *testing.T) {
//	fmt.Println("测试添加用户：")
//	user := &User{}
//	user.AddUser()
//	user.AddUser2()
//}

//func TestUserGetUserById(t *testing.T) {
//	fmt.Println("测试查询一条记录：")
//	user := User{
//		ID:       1,
//	}
//	u, _ := user.GetUserById()
//	fmt.Println("得到的user信息是：", u)
//}

func TestUserGetUsers(t *testing.T) {

	fmt.Println("测试查询所有的记录：")
	user := User{}
	u, _ := user.GetUsers()
	//遍历user切片
	for k, v := range u {
		fmt.Printf("打印第%v个用户是%v: \n", k+1, v)
	}

}
