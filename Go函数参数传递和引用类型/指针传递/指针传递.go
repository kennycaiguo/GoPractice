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
	changename(&jim)  //传进函数的是变量jim的指针
	fmt.Println(jim) //{another 10}
}

func changename(p *person){
	p.name = "another" //这里是指针指向的变量赋值，所以有效果
	/*
		细心的同学可能发现p是一个指针变量，为什么能直接加.name?
		此处是go语言编译器自动转换了写法
		准确来写应该是：
		(*p).name = "another"
	*/
}