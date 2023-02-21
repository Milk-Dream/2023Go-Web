package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

//1.手机卡
const ipAndPort = "127.0.0.1:9090"

func main() {
	//2.插手机卡，绑定电话卡
	listen, _ := net.Listen("tcp", ipAndPort)
	fmt.Println("电话开机了，等待来电了...")
	//等电话
	conn, _ := listen.Accept()
	fmt.Println("电话来了")
	reader := bufio.NewReader(conn)
	msg := make([]byte, 100)
	reader.Read(msg)
	msgStr := string(msg)
	fmt.Println("收到:", msgStr)
	//4.回消息
	conn.Write([]byte(strings.ToUpper(msgStr)))
	//挂电话
	conn.Close()
}