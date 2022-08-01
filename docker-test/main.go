package main

import (
	"fmt"
	"html"
	"net/http"
)

type TaskAuditState int

const (
	TaskWaitHandle   TaskAuditState = iota + 1 // 待审核
	TaskSecondReview                           // 复审
	TaskPass                                   // 通过
	TaskRefuse                                 // 拒绝
)

func main() {
	fmt.Println(TaskWaitHandle, TaskRefuse)

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "hello, %q", html.EscapeString(request.URL.Path))
	})
	fmt.Println("hello docker, server is running...")
	http.ListenAndServe(":9999", nil)

}
