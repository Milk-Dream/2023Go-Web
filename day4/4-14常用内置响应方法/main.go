package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe("127.0.0.1:8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	if method == "GET" {
		// http.NotFound(w, r)
		// http.Redirect(w, r , "https://www.baidu.com/", 302)
		// http.Redirect(w, r , "https://www.baidu.com/", http.StatusFound)
		http.ServeFile(w, r, r.URL.Path[1:])
		return
	}

	fmt.Fprintln(w, "helloWorld")
}