package main

import "net/http"

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe("127.0.0.1:8080", nil)
}


func index(w http.ResponseWriter, r *http.Request) {
	//Write，响应头没有指定Content-Type，则写入的前512字节会被用于检测Content-Type
	w.Write([]byte("hello world"))
}