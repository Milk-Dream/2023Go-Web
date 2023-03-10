package routes

import "net/http"

func registerHomeRoutes() {
	http.HandleFunc("/home", home)
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("home!"))
}
