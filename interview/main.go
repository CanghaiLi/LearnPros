package main

import (
	"fmt"
	"time"
	"unicode/utf8"
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
	fmt.Println(len("he1, 啊"))
	fmt.Println(utf8.RuneCountInString("he1, 啊"))

	//time.Sleep(time.Second * 1)
	println("loading")

	for c := range ch {
		println(c)
	}
	fmt.Println("exit")
}
