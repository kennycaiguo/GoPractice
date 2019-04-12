package main
import (
	"fmt"
)
func main(){
	// 创建一个整型切片
	// 其容量和长度都是5个元素
	slice := []int{10, 20, 30, 40, 50}
	// 创建一个新切片
	// 其长度为2个元素，容量为4个元素
	newSlice := slice[1:3]

	fmt.Println(slice)
	// [10 20 30 40 50]
	fmt.Println(newSlice)
	// [20 30]
}

