package main

import (
	"log"
	"net/http"
)

func main() {

	log.Fatalln(http.ListenAndServe(":8888", nil))
}
