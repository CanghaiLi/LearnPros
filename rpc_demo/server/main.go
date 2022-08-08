package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
)

// World rcvr类对象
type World struct {
}

func (w World) Helloworld(input string, output *string) error {
	*output = input + " 你好！"
	return nil
}

func main() {
	//	1. 注册rpc服务，绑定方法
	//rpc.RegisterName("Demo", new(World))
	err := rpc.Register(new(World))
	if err != nil {
		log.Fatalln(err, "注册rpc服务失败")
	}

	//	2. 设置监听
	fmt.Println("开始监听...")
	listener, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Fatalln(err, "设置监听服务失败")
	}
	defer listener.Close()

	for {
		//	3. 建立链接
		conn, err := listener.Accept()
		fmt.Println("建立链接...")
		if err != nil {
			fmt.Println(err, "建立链接务失败")
			continue
		}

		//	4. 绑定服务
		go rpc.ServeConn(conn)
	}

}
