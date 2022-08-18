package main

import (
	_ "embed"
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

type Js struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

//go:embed file.txt
var s string

func main() {
	fmt.Println(s)
	//bytes, err := json.Marshal(Js{"lee", 18})
	//if err != nil {
	//	return
	//}
	//qrcode.WriteFile(string(bytes), qrcode.Medium, 256, "./golang_qrcode.png")
	num := 123.1
	f2, ok := interface{}(num).(float64)
	fmt.Println(f2, ok)
}
