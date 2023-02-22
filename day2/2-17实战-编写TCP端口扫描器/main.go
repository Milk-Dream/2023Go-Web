package main

import (
	"fmt"
	"net"
	"sync"
)

func scanner(addr string, i int) {
	defer wg.Done()
	
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		fmt.Println(addr, "不能使用")
		return
	}
	fmt.Println(addr, "可以使用")
	if conn != nil {
		conn.Close()
	}
	
	
	
}

var wg sync.WaitGroup

func main() {
	var (
		ipAddr string
		bnum int
		enum int
	)
	fmt.Println("==========欢迎使用小黑老师s端口扫描器==========")
	fmt.Print("请输入要扫描的IP地址:")
	fmt.Scanln(&ipAddr)
	fmt.Print("请输入扫描的范围(开始:begin 1~65534)")
	fmt.Scanln(&bnum)
	fmt.Print("请输入扫描的范围(结束:end 1~65535)，结束的范围要比开始大")
	fmt.Scanln(&enum)
	fmt.Println("开始扫描")
	for i := int(bnum); i <= int(enum); i++ {
		wg.Add(1)
		addr := fmt.Sprintf("%s:%d", ipAddr,i)
		go scanner(addr, i)
	}
	wg.Wait()
	fmt.Println("扫描结束")
}