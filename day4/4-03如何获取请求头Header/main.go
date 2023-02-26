package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	for k, v := range(r.Header) {
		fmt.Fprintln(w, k, v)
	}
	fmt.Fprintln(w, "-------------")
	fmt.Fprintln(w, r.Header["User-Agent"])
	fmt.Fprintln(w, r.Header["User-Agent"][0])
	//Get 获取第一个参数值，如果不存在返回空字符串
	fmt.Fprintln(w, r.Header.Get("User-Agent"))
}