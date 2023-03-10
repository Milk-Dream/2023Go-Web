package routes

import "net/http"

func registerIndexRoutes() {
	http.HandleFunc("index", IndexHandler)
}


func IndexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("index!"))
}