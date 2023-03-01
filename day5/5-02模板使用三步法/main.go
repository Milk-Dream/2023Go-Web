package main

import (
	"fmt"
	"net/http"
	"text/template"
)

func main() {
	http.HandleFunc("/baidu", baidu)
	http.HandleFunc("/jd", jd)
	http.HandleFunc("/tx", tx)
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
	tmp.Execute(w, "小黑")
	//http.ServeFile(w, r, "baidu.html")
}
func jd(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "jd.html")
}
func tx(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "tx.html")
}