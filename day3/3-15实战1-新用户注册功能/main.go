package main

import (
	"fmt"
	"net/http"
)

var users = make(map[string]string, 100)

func register(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	username := r.Form.Get("username")
	password := r.Form.Get("password")
	re_password := r.Form.Get("re_password")

	//判断密码是否输入一致
	if (re_password != password) {
		w.Write([]byte("密码不一致"))
		return
	}

	//判断用户名是否存在
	_, ok := users[username]
	if (ok) {
		w.Write([]byte("用户名已经存在"))
		return
	}


	users[username] = password
	w.Write([]byte(username+"注册成功"))

}

func main() {

	http.HandleFunc("/register", register)
	fmt.Println("服务端开启中...")
	http.ListenAndServe(":8080", nil)
}