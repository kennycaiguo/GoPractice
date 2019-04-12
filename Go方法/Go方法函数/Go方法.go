package main

import (
	"fmt"
)

type user struct {
	name  string
	email string
}
//传递值，changeEmail方法没有用
func (u user) print() {
	fmt.Println(u.email)
}
func (u user) changeEmail(email string) {
	u.email = email
}
func main() {
	bill := &user{"Bill", "bill@email.com"}
	bill.print()
	//bill@email.com
	bill.changeEmail("new address")
	bill.print()
	//bill@email.com
}
