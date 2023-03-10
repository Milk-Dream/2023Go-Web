package main

import "net/http"

type AuthMiddleWara struct {
	Next http.Handler
}

func (mm *AuthMiddleWara) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if mm.Next == nil {
		mm.Next = http.DefaultServeMux
	}

	//校验工作
	if r.Header.Get("X-Auth") != "" {
		w.Write([]byte("校验通过"))
		mm.Next.ServeHTTP(w, r)
	} else {
		http.NotFound(w, r)
		return
	}
}

func main() {
	http.ListenAndServe("127.0.0.1:8080", new(AuthMiddleWara))
}