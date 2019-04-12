package main

import (
	"fmt"
)

func main(){
	array :=[5]*int{}
	array[0]=new(int) //
	*array[0]=1
	fmt.Printf("%d",*array[0]) // 1
}

