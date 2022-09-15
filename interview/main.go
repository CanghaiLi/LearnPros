package main

import (
	"crypto/md5"
	_ "embed"
	"fmt"
	"io"
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

type Stud struct {
	Nums string
}

type Js struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Stud
}

//go:embed file.txt
var s string

func main() {
	fmt.Println(s)
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	fmt.Println("4.特定格式时间: ", time.Now().Format("15:04:05"))
	js := Js{
		Name: "",
		Age:  0,
		Stud: Stud{
			Nums: "1,2,3,4",
		},
	}
	fmt.Println(js.Nums)
	fmt.Println(js.Stud.Nums)

	bytes := Md5Demo("fullstack!fullbook")

	fmt.Printf("%x\n", bytes)
	Md5HashDemo("fullstack!", "fullbook")
}

func Md5Demo(s string) [16]byte {
	data := []byte(s)
	return md5.Sum(data)
}

// Md5HashDemo 返回一个新的使用MD5校验的hash.Hash接口。
func Md5HashDemo(s, s1 string) {
	h := md5.New()
	io.WriteString(h, s)
	io.WriteString(h, s1)
	fmt.Printf("%x\n", h.Sum(nil))
}
