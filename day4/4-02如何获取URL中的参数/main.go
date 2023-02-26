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
	//URL.RawQuery
	//URL.Query()

	fmt.Println(r.URL.RawQuery)
	fmt.Println(r.URL.Query())
	fmt.Fprintln(w, r.URL.RawQuery)
	fmt.Fprintln(w, r.URL.Query())
	fmt.Fprintln(w, r.URL.Query()["name"][0])
	fmt.Fprint(w, r.URL.Query().Get("name"))//获取第一个参数，如果没有则返回空字符串

	//www.baidu.com/123/456
	// 123/456
	fmt.Fprintln(w, r.URL.Path)
	//123/456
	r.URL.EscapedPath()
	//课程没讲清楚
	fmt.Fprintln(w, r.URL.RawPath)

}