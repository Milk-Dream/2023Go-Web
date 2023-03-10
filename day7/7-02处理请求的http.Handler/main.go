package main

import (
	"fmt"
	"net/http"
)

type MyHandler struct {}

type database map[string]int

func (db database) ServeHTTP (w http.ResponseWriter, r *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %d\n", item, price)
	}
}
/*
func (mh *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("I am Dog"))
}
*/
func main() {
	//mh := &MyHandler{}

	//http.ListenAndServe("localhost:8080", mh)
	//http.ListenAndServe("localhost:8080", http.DefaultServeMux)
	/*
	mh := &MyHandler{}
	server := http.Server {
		Addr: "localhost:8080",
		Handler: mh,
	}
	server.ListenAndServe()
	*/
	db := database{"apple":5, "pear":6, "orange": 7}
	server := http.Server {
		Addr: "localhost:8080",
		Handler: db,
	}
	server.ListenAndServe()
}