package main

import (
	"fmt"
	"runtime"
)

func say(s string) {
	for i := 0; i < 2; i++ {
		runtime.Gosched()
		fmt.Println(s," i:",i)
	}
}

func main() {
	runtime.GOMAXPROCS(1)
	go say("world")
	say("hello")
}
