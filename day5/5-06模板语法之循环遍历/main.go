package main

import (
	"fmt"
	"net/http"
	"text/template"
)

type User struct {
	Name string
	Age int
}

func main() {
	http.HandleFunc("/baidu", baidu)
	
	http.ListenAndServe("127.0.0.1:8080", nil)
}

func baidu(w http.ResponseWriter, r *http.Request) {
	//编写百度
	//1.编写模板
	//2.解析模板
	tmp, err := template.ParseFiles("baidu.html")
	if err != nil {
		fmt.Println("模板解析失败:", err)
		return
	}

	//渲染模板
	//user := User{"jack", 18}
	books := map[string]string{"book1":"西游记", "book2":"红楼梦"}

	//其中的.就是这个
	//tmp.Execute(w, user)
	tmp.Execute(w, books)
	//http.ServeFile(w, r, "baidu.html")
}
