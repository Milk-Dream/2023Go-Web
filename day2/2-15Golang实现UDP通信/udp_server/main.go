package main

import (
	"fmt"
	"net"
)

func main() {
	listen, _ := net.ListenUDP(
		"udp",
		&net.UDPAddr{
			IP:net.IPv4(0, 0, 0, 0),
			Port:9090,
		},

	)
	//收数据
	var data [1024]byte
	n, addr, _ := listen.ReadFromUDP(data[:])
	fmt.Printf("data:%v, addr:%v, counts: %v\n", string(data[:n]), addr, n)
	//发数据
	listen.WriteToUDP(data[:n], addr)
}