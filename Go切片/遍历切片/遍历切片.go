package main
import (
	"fmt"
)
func main(){
	// 创建一个整型切片
	// 其长度和容量都是4个元素
	slice := []int{10, 20, 30, 40}
	// 迭代每一个元素，并显示其值
	for index, value := range slice {
	fmt.Printf("Index: %d　Value: %d\n", index, value)
	}
	for i:=0;i<len(slice);i++ {
		fmt.Println("value:",slice[i])
	}
}

