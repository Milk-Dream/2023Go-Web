package main

import "net/http"



func main() {
	/*http.HandleFunc("/index", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("index"))
	})
	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("about"))
	})
	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("home"))
	})*/
	routes.RegisterRoutes()





	http.ListenAndServe("localhost:8080", nil)

	
}