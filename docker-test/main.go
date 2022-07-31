package main

import (
	"fmt"
	"html"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "hello, %q", html.EscapeString(request.URL.Path))
	})
	fmt.Println("hello docker, server is running...")
	http.ListenAndServe(":9999", nil)

}
