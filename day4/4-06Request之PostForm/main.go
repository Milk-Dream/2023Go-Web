package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request)  {
	//PostForm字段,不会获取url中的数据
	//www.baidu.com/?username=1529952847
	//username = 188
	//PostForm和form只支持application/x-www-form-urlencode
	r.ParseForm()
	fmt.Fprintln(w, r.PostForm)
	fmt.Fprintln(w, r.Form)


}