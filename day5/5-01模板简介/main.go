package main

import "net/http"

func main() {
	http.HandleFunc("/baidu", baidu)
	http.HandleFunc("/jd", jd)
	http.HandleFunc("/tx", tx)
	http.ListenAndServe("127.0.0.1:8080", nil)
}

func baidu(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "baidu.html")
}
func jd(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "jd.html")
}
func tx(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "tx.html")
}