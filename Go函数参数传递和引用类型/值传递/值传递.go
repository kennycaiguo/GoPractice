package main

import(
	"fmt"
)

type person struct{
	name string
	age int
}

func main(){
	jim := person{"jim",10}
	fmt.Println(jim) //{jim 10}
	changename(jim)  //传进函数的是值的拷贝
	fmt.Println(jim) //{jim 10}
}

func changename(p person){
	p.name = "another" //这里是对拷贝赋值，没有效果
}