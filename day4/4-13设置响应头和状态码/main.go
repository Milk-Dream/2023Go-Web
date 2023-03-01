package main

import "net/http"

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe("127.0.0.1:8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("x-YYY", "my_self")
	w.Header().Set("x-YYY", "your_self")

	//writer之后，响应头就没办法修改和设置
	w.WriteHeader(400)

	w.Write([]byte("hello"))
	//Write之后修改是没有效果的
	// w.Header().Add("x-YYY", "my_self")
	// w.Header().Set("x-YYY", "your_self")

}