package routes

import "net/http"

func registerAboutRoutes() {
	http.HandleFunc("about", AboutHandler)
	
}

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("about!"))
}