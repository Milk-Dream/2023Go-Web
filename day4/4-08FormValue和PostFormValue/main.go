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
	//FormValue无需调用ParseMultipartForm
	//只能获取第一个值，并且是普通字段
	//当表单类型是urlencoded时候，form中表单字段在前，返回表单中的值
	//当表单类型是form-data时，Form中有URL的值，则返回URL中的值，如果URL没有值，则在表单找
	
	fmt.Fprintln(w, r.FormValue("username"))
	
	// fmt.Fprintln(w, r.FormValue("avatar"))
	fmt.Fprintln(w, r.FormValue("time"))

	//PostForm只返回表单中的数据
	fmt.Fprintln(w, r.PostForm)
	//只能获取普通字段
	fmt.Fprintln(w, r.PostForm.Get("username"))
	fmt.Fprintln(w, r.PostForm.Get("avatar"))
}