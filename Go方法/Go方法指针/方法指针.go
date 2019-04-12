package main

import (
   "fmt"
)

// user在程序里定义一个用户类型
type user struct {
	name string
    email string
}

// 函数接收user类型的指针
func (u *user) printPointer() {
	fmt.Println(u.email)
}
func (u *user) changeEmailPointer(email string) {
	u.email = email
}

func main() {
	bill := &user{"Bill", "bill@email.com"}
	bill.printPointer() 
	//bill@email.com
	bill.changeEmailPointer("new address") 
	bill.printPointer()
	//new address
}