package main

import (
	"sync"
	"runtime"
	"fmt"
)
func main(){
	//调度器只能为该程序使用一个逻辑处理器
	runtime.GOMAXPROCS(1)

	var wg sync.WaitGroup
	wg.Add(2)

	fmt.Println("start goroutine")


	go func (){
		defer wg.Done()
		for outer := 2; outer < 5000; outer++ {
			for inner := 2; inner < outer; inner++ {
				if outer%inner == 0 {
					continue
				}
			}
			fmt.Printf("A:%d\n", outer)
		}
	}()

	go func (){
		defer wg.Done()
		for outer := 2; outer < 5000; outer++ {
			for inner := 2; inner < outer; inner++ {
				if outer%inner == 0 {
					continue
				}
			}
			fmt.Printf("B:%d\n", outer)
		}
	}()

	fmt.Println("waiting to finish")

	wg.Wait()

	fmt.Println("Terminating Program")
}