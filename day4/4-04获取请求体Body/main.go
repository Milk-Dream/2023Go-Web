package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("服务端开启中...")
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	//获取请求体参数
	body := make([]byte, 100)
	n, _ := r.Body.Read(body)
	fmt.Fprintln(w, n)
	fmt.Fprintln(w, "请求体中的数据", n, string(body))


}