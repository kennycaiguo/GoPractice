// Sample program to show how polymorphic behavior with interfaces.
package main

import (
	"fmt"
)

type person interface {
	sayInfo()
}

type user struct {
	name  string
	email string
}

type admin struct {
	name  string
	email string
}

func (u *user) sayInfo() {
	fmt.Printf("Sending user email to %s<%s>\n",
		u.name,
		u.email)
}

func (a *admin) sayInfo() {
	fmt.Printf("Sending admin email to %s<%s>\n",
		a.name,
		a.email)
}

func main() {
	bill := user{"Bill", "bill@email.com"}
	sendInfo(&bill)

	lisa := admin{"Lisa", "lisa@email.com"}
	sendInfo(&lisa)
}

func sendInfo(p person) {
	p.sayInfo()
}
