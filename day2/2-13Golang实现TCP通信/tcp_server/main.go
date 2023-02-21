package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

//1.手机卡
const ipAndPort = "127.0.0.1:9090"

func process(conn net.Conn) {
	fmt.Println("电话来了")
	//挂电话，不要忘记关闭连接
	reader := bufio.NewReader(conn)
	defer conn.Close()
	for {
		msg, err:= reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				log.Println("connection closed")
			} else {
				log.Println("connection closed")
			}
			
		}
		
		msgStr := string(msg)
		fmt.Println("收到:", msgStr)
		//4.回消息
		conn.Write([]byte(strings.ToUpper(msgStr)))
		
	}
	
	
}

func main() {
	//2.插手机卡，绑定电话卡
	listen, _ := net.Listen("tcp", ipAndPort)
	fmt.Println("电话开机了，等待来电了...")
	
	for {
		//等电话
		conn, _ := listen.Accept()
		process(conn)
	}
	
	
}