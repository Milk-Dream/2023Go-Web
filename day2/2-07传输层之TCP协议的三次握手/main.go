package main

func main() {
	//TCP协议是好人协议:只要有请求就有回复
	//TCP协议是可靠协议，TCP请求确认机制保证了数据安全
	//三次握手:建立链接，保证链接可靠
	//第一次:客户端主动朝服务端发送请求，希望建立链接
	//第二次:服务端收到请求后回复同意；然后服务端向客户端请求建立链接。两次并做一次
	//第三次:客户端收到请求，回复同意
	//整个过程就是三次握手建立TCP连接
}