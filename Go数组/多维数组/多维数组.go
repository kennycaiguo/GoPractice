package main

import (
	"fmt"
)

func main(){
	array :=[5][3]int{0:{11,22,33}, 1:{44,55,66}}
	/*array为[
		[11,22,33],
		[44,55,66],
		[0,0,0],
		[0,0,0],
		[0,0,0],
	]*/
	for i:=0;i<5;i++{
		for j:=0;j<3;j++{
			fmt.Printf("[%d][%d]=%d\n",i,j,array[i][j])
		}
	}
	fmt.Printf("\n")//或者使用range循环
	for i,v1:=range array{
		for j,v2:= range v1{
			fmt.Printf("[%d][%d]=%d\n",i,j,v2)
		}
	}

}

