package main

import (
	"fmt"
	"net/rpc"
)

func main() {
	// 1. rpc链接服务器
	conn, err := rpc.Dial("tcp", "localhost:8888")
	if err != nil {
		return
	}

	//	2. 远程调用
	var res string
	conn.Call("World.Helloworld", "Lining", &res)
	fmt.Println(res)
}
