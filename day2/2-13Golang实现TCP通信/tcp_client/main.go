package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

//1.手机卡
const ipAndPort = "127.0.0.1:9090"

func main() {
	//1.拨电话
	conn, _ := net.Dial("tcp", ipAndPort)
	//2.发消息
	reader := bufio.NewReader(os.Stdin)
	defer conn.Close()
	for {
		input, _ := reader.ReadString('\n')
		conn.Write([]byte(input))
		//3.收消息
		buf := make([]byte, 100)
		conn.Read(buf)
		fmt.Println("收到:",string(buf))
	}
	
	
	

}