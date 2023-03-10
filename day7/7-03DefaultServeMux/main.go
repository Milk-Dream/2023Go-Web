package main

import (
	"fmt"
	"net/http"
	
)

type database map[string]int

func (db database) ServeHTTP(w http.ResponseWriter, r http.Request) {
	path := r.URL.Path
	if path == "/" {
		for item, price := range db {
			fmt.Fprintf(w ,"%s,%s", item, price)
		}
	} else {
		w.Write([]byte("你不是根路径"))
	}
}

func main() {
	db := database{"apple": 5, "pear": 6, "orange": 10}
	fmt.Println(db)
	server := http.Server {
		Addr: "localhost:8080",
		Handler: nil, //http.DefaultServerMux
	}
	server.ListenAndServe()

	
}