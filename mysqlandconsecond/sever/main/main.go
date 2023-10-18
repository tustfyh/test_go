package main

import (
	_ "encoding/binary"
	"fmt"
	"mysqlandcon/sever/model"
	"net"
	"time"
)

// 服务器
func process(conn net.Conn) {
	//接受客户端发送的消息
	//defer conn.Close()
	processor := &Processor{Conn: conn}
	err := processor.process02()
	if err != nil {
		return
	}
}
func init() {
	//初始化连接池
	initPool("127.0.0.1:6379", 16, 0, 300*time.Second)
	//初始化userdao
	model.MyuserDao = model.NewRedisConn(Pool)
	//初始化在线列表

}
func main() {
	fmt.Println("服务器开始监听")
	listen, err := net.Listen("tcp", "127.0.0.1:9000")
	if err != nil {
		return
	}
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Printf("listen.Accept() err=%v\n", err)
		}
		go process(conn)
	}

}
