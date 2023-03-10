package routes

import "net/http"

func registerAboutoutes() {
	http.HandleFunc("/about", about)
}

func about(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("About!"))
}
