package main
import (
	"fmt"
)
func main(){
	// 创建一个整型切片
	// 其容量和长度都是5个元素
	slice := []int{10, 20, 30, 40, 50}
	// 改变索引为1的元素的值
	slice[1] = 25
	fmt.Println(slice)
	// [10 25 30 40 50]
}

