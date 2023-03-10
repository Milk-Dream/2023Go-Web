package main

import (
	"1_server/routes"
	"log"
	"net/http"
)

type AuthMiddleWare struct {
	Next http.Handler
}

func (mm *AuthMiddleWare) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if mm.Next == nil {
		mm.Next = http.DefaultServeMux
	}
	// 校验工作
	if r.Header.Get("X-Auth") != "" {
		mm.Next.ServeHTTP(w, r)
	} else {
		http.NotFound(w, r)
		return
	}
}

func main() {
	// 设置类的工作
	routes.RegisterRoutes()

	log.Fatal(http.ListenAndServe("localhost:8080", new(AuthMiddleWare)))
}
