package routes

import "net/http"

func registerIndexRoutes() {
	http.HandleFunc("/index", index)
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("index!"))
}
