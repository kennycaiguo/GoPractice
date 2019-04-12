package main
import (
	"fmt"
)
func main(){
	source := []string{"Apple", "Orange", "Plum", "Banana", "Grape"}
	// 对第三个元素做切片，并限制容量
	// 其长度和容量都是1个元素
	slice := source[2:3:3]
	// 向slice追加新字符串
	slice = append(slice, "Kiwi")
	slice[0]="newItem"
	fmt.Println(source)
	// [Apple Orange Plum Banana Grape]
	fmt.Println(slice)
	// [newItem Kiwi]
}

