package main

import "net/http"

func main() {
	
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("HelloWorld"))

	})
	
	http.ListenAndServeTLS("localhost:8080","cert.pem", "key.pem", nil)

	//go run D:\Go\src\crypto\tls\generate_cert.go -host localhost
}