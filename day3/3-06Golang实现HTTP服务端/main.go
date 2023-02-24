package main

import (
	"fmt"
	
	"net/http"
)

func get(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("获取资源成功"))
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>index 页面</h1>"))
}

func main() {
	//注册一个路由，当用户请求/get，就会由get函数处理
	http.HandleFunc("/get", get)
	http.HandleFunc("/index", index)
	http.HandleFunc("/set", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("设置成功"))
	})
	fmt.Println("服务端启动中...")
	//使用nil的话，就会默认使用http.DefaultServeMux
	//ListenAndServer负责开启服务
	http.ListenAndServe("127.0.0.1:80", nil)
	//127.0.0.1是可以简化的
	//http.ListenAndServe(":80", nil)
}