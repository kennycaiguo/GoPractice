package main

import (
	"fmt"
)

func main(){
	array :=[...]int{0:1,1:1,4:0}
	//array为[1,1,0,0,0]
	for i:=0;i<5;i++{
		fmt.Printf("%d ",array[i])
		//1 1 0 0 0
	}
	fmt.Printf("\n")
	for index,value:= range array{
		fmt.Printf("index:%d,value:%d\n",index,value)
	}
}

