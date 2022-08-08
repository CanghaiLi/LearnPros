package main

import (
	"fmt"
	"time"
)

var ch = make(chan int, 5)

func f() {
	time.Sleep(time.Second * 1)
	println("start")
	for i := 1; i < 5; i++ {
		ch <- i
	}
	close(ch)
	println("close")
}

func main() {
	go f()

	//time.Sleep(time.Second * 1)
	println("loading")

	for c := range ch {
		println(c)
	}
	fmt.Println("exit")
}
