package main

import "net/http"

func main() {

	//第一种启动server
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world"))
	})
	http.ListenAndServe("localhost:8080", nil)
	//第二种启动server
	server := http.Server{
		Addr: "localhost:8080",
		Handler: nil,
		
	}
	server.ListenAndServe()

	
}