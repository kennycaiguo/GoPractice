package main

import (
	"fmt"
)


type printInfor interface {
	printInfo()
}

type user struct {
	name  string
	email string
}

func (u *user) printInfo() {
	fmt.Println("name:", u.name, " email:",u.email)
}

func main() {
	u := user{"Bill", "bill@email.com"}
	sendInfo(&u)
}

func sendInfo(p printInfor) {
	p.printInfo()
}
