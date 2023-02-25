package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Proto", r.Proto)
	fmt.Println("Method", r.Method)
	fmt.Println("User-Agent", r.UserAgent())
	
}

func main() {
	//r http.Request
	//Cookies
	//FormValue
	//Referer
	//UserAgent
	http.HandleFunc("/", index)
	fmt.Println("服务端开启中...")
	http.ListenAndServe(":8080", nil)

}