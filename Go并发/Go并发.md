#### Go并发

##### 进程和线程


##### goroutine实际上是协程

什么是协程？

在**一个线程**中,假设有2个子程序A,B，作用是求5000以内的素数，并打印
```go
func A(){
    for outer := 2; outer < 5000; outer++ {
        for inner := 2; inner < outer; inner++ {
            if outer%inner == 0 {
                continue
            }
        }
        fmt.Printf("A:%d\n", outer)
    }
    fmt.Printf("A子程序结束")
}
func B(){
    for outer := 2; outer < 5000; outer++ {
        for inner := 2; inner < outer; inner++ {
            if outer%inner == 0 {
                continue
            }
        }
        fmt.Printf("B:%d\n", outer)
    }
    fmt.Printf("B子程序结束")
}
```
如果这两个子程序是以协程的形式调用，打印的结果会是怎样的?

输出
```js
> B:2~B:3467     //运行一段时间子程序B，子程序B挂起，开始运行子程序A
> A:2~A:3467     //运行一段时间子程序A，子程序A挂起，开始运行子程序B
> B:3469~B:4999  //运行一段时间子程序B，子程序B挂起，开始运行子程序A
> B子程序结束     // B子程序结束
> A:3469~A:4999  //运行一段时间子程序A
> A子程序结束     // A子程序结束
```

协程就是在一个线程中多个子程序切换运行，这样就叫做协程。

而且注意到上例中协程的切换是程序内部自己调度的。上述例子的完整代码如下，大家可以自己尝试：
```go
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
```

协程的好处？

1.相对于多线程，协程可以节省线程切换的开销，协程是程序内部自己调度，性能更好

2.不需要多线程的锁机制，因为只有一个线程，也不存在同时写变量冲突，在协程中控制共享资源不加锁，只需要判断状态就好了，所以执行效率比多线程高很多。









