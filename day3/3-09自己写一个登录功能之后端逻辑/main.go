package main

import (
	"fmt"
	"net/http"
)

func login(w http.ResponseWriter, r *http.Request) {
	//1.接收前端传过来的参数:用户名，密码
	username := r.URL.Query().Get("username")
	password := r.URL.Query().Get("password")
	//2.校验用户名和密码是否一致
	if(username == "Leo" && password == "123456") {
		w.Write([]byte("<h1>Leo登录成功</h1>"))
	} else {
		w.Write([]byte("<h1>登录失败，请检查你的用户名和密码</h1>"))
	}
}

func main() {

	fmt.Println("服务的开启中...")
	http.HandleFunc("/login", login)
	http.ListenAndServe(":8080", nil)
}