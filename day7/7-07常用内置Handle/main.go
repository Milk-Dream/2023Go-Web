package main

import (
	"net/http"
	"time"
)

type myHandler struct{}

func (mh myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	time.Sleep(5 * time.Second)
	w.Write([]byte("hello World"))
}

func main() {
	http.Handle("/apple", http.NotFoundHandler())
	http.Handle("/orange", http.RedirectHandler("http://www.baidu.com", 307))
	http.Handle("/sleep", http.TimeoutHandler(&myHandler{},2* time.Second,"超时了"))
	http.ListenAndServe("127.0.0.1:8080", nil)
}