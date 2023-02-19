package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	// w.Write([]byte("hello world"))
	// s := "<h1 style='color:red'>Hello World</h1>"
	s, _:= ioutil.ReadFile("./test.txt")
	// w.Write([]byte(s))
	w.Write(s)
}

func main() {
	http.HandleFunc("/", hello)
	//此时代码会等待，等待请求
	http.ListenAndServe("127.0.0.1:8080", nil)
	fmt.Println("服务开启成功，请访问127.0.0.1:8080")
}