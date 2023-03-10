package routes

import "net/http"

func registerHomeRoutes() {
	http.HandleFunc("home", HomeHandler)
	
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("home!"))
}