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
	//Form字段
	r.ParseForm() //先解析，再使用
	fmt.Fprintln(w, r.Form)//map[string]stirng
	fmt.Fprintln(w, r.Form["username"][0])
	fmt.Fprintln(w, r.Form.Get("username"))

}