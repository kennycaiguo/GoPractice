package main

import(
	"fmt"
)

type user struct{
	name string
	age int
	array []int
}

func main(){
	jim := user{"jim",10,[]int{1,2}}
	fmt.Println(jim)
	changename(jim)
	fmt.Println(jim)
}

func changename(p user){
	p.name = "another"
	p.array[0] = 33333
}