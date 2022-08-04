package main

import (
	"database/sql"
	"fmt"
)

type TaskAuditState int

const (
	TaskWaitHandle   TaskAuditState = iota + 1 // 待审核
	TaskSecondReview                           // 复审
	TaskPass                                   // 通过
	TaskRefuse                                 // 拒绝
)

type Student struct {
	name string
	Person
}
type Person struct {
	IdCard int
}

func (p *Person) sayHello() {
	fmt.Println("hello")
}
func (p Person) setID(id int) {
	p.IdCard = id
}

type Op func(msg any) (any, error)

func Foo(param any) Op {

	return func(msg any) (any, error) {
		return msg, nil
	}
}

var DB *sql.DB

func main() {
	fmt.Println(TaskWaitHandle, TaskRefuse)
	var p0 = Student{}
	p := &Person{
		IdCard: 123,
	}
	p0.sayHello()
	p.setID(321)
	fmt.Println(p)

	//http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
	//	fmt.Fprintf(writer, "hello, %q", html.EscapeString(request.URL.Path))
	//})
	//fmt.Println("hello docker, server is running...")
	//http.ListenAndServe(":9999", nil)

}
