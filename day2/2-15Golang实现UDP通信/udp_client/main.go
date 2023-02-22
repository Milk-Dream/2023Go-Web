package main

import (
	"fmt"
	"net"
)

func main() {
	socket, _ := net.DialUDP(
		"udp",
		nil,
		&net.UDPAddr{
			IP: net.IPv4(0, 0, 0, 0),
			Port:9090,
		},
	)
	defer socket.Close()
	//发送数据
	socket.Write([]byte("hello World"))
	//接收数据
	data := make([]byte, 4096)
	n, addr, _ := socket.ReadFromUDP(data)
	fmt.Printf("recv:%v,addr:%v,counts,%v",string(data[:n]),addr,n)
}