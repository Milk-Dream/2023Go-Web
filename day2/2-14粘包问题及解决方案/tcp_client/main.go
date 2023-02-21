package main

import (
	"fmt"
	"net"
)

func main() {
	//粘包的原因
	//发送端:单次信息量小，操作系统批量发出去的
	//接收端:本次没有接收完，剩下的部分粘在了一起

	//粘包的解决方案
	//告诉对方我的消息长度是多少
	conn, _ := net.Dial("tcp", "127.0.0.1:9090")
	defer conn.Close()
	//注意:循环多次发数据，单次信息少，多次一块发送了
	for i:=0; i < 10; i++ {
		msg := "Are you Ok"
		conn.Write([]byte(msg))
	}

	buf := make([]byte, 100)
	conn.Read(buf)
	fmt.Println("收到:", string(buf))
}