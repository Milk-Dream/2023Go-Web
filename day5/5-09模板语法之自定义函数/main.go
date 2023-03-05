package main

import (
	
	"net/http"
	"text/template"
)

type User struct {
	Name string
	Age int
}

type Student struct {
	Name string
	Age int
	Scores map[string]int
}

func info(w http.ResponseWriter, r *http.Request) {
	add := func(a, b int)(int,error) {
		return a + b, nil
	}

	//自定义函数，返回值只可以一个值，或者一个值和一个错误
	tmp, _ := template.New("test.html").Funcs(template.FuncMap{"add":add}).ParseFiles("test.html")

	//链式调用
	/*
	tmp, _ = tmp.ParseFiles("test.html")
	*/
	tmp.Execute(w, nil)


}


func main() {
	http.HandleFunc("/test", info)
	
	http.ListenAndServe("127.0.0.1:8080", nil)
}

