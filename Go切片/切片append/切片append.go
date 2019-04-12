package main
import (
	"fmt"
)
func main(){
	// 创建一个整型切片
	// 其长度和容量都是4个元素
	slice := []int{10, 20, 30, 40}
	// 向切片追加一个新元素
	// 将新元素赋值为50
	newSlice := append(slice, 50)
	// newSlice的length为5，cap为8
	fmt.Println(newSlice)
}

