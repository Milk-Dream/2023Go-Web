package main

import (
	"fmt"
	"net/http"
)

func main() {
	//MultipartForm类型
	//第一个字段 存表单中的普通数据
	//第二个字段 FileHeader (存表单提交的文件数据)
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}


func index(w http.ResponseWriter, r *http.Request) {
	//formdata
	r.ParseMultipartForm(1024)
	fmt.Fprintln(w, r.MultipartForm)
	fmt.Fprintln(w, r.MultipartForm.Value)
	fmt.Fprintln(w, r.MultipartForm.File["avatar"][0])
}