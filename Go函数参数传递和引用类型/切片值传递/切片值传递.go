package main

import(
	"fmt"
)

func main(){
	persons := []string{"jim","tom","panda"}
	fmt.Println(persons) //[jim tom panda]
	changeperson(persons)  //传进函数的是值的拷贝
	fmt.Println(persons) //[another tom panda]
}

func changeperson(p []string){
	p[0] = "another" //这里是对拷贝赋值，没有效果
}