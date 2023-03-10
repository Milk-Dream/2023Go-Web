package main

import (
	"log"
	"net/http"
)

type MyMiddleWara struct {
	Next http.Handler

}

func (mm *MyMiddleWara)ServeHTTP (w http.ResponseWriter, r *http.Request) {
	if mm.Next == nil {
		mm.Next = http.DefaultServeMux
	}
	//在请求之前做一个操作
	log.Println(111)
	mm.Next.ServeHTTP(w, r)
	w.Write([]byte("hello World"))
}

func main() {
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", new(MyMiddleWara)))
}