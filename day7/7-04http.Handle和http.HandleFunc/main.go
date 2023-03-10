package main

import (
	
	"net/http"
	
)


//http.Handle
//http.HandleFunc

type RootHandler struct{}
type IndexHandler struct{}
func (mh *RootHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("this is /"))
}

func (ih *IndexHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("this is /index"))
}

func about(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("this is about"))
}

func main() {
	http.Handle("/", &RootHandler{})
	http.Handle("/index", &IndexHandler{})
	
	http.HandleFunc("/about", about)

	http.ListenAndServe("localhost:8080", nil)
	
}