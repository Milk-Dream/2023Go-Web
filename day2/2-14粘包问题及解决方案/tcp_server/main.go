package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

func process(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	var buf [50]byte //注意:每次接收50个字节数
	for {
		n, err := reader.Read(buf[:])
		if err != nil {
			if err == io.EOF {
				log.Println("connection close")
			} else {
				log.Println(err)
			}
			return
		}
		msg := string(buf[:n])
		fmt.Println("收到:", msg)
		conn.Write([]byte(strings.ToUpper(msg)))
	}
}

func main() {
	//粘包的原因
	//发送端:单次信息量小，操作系统批量发出去的
	//接收端:本次没有接收完，剩下的部分粘在了一起
	
	//粘包的解决方案
	//告诉对方我的消息长度是多少
	listen, _ := net.Listen("tcp", "127.0.0.1:9090")
	defer listen.Close()
	fmt.Println("电话开机了，等待来电中...")
	for {
		conn, _ := listen.Accept()
		fmt.Println("电话来了")
		go process(conn)
	}
}