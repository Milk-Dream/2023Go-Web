package routes

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

var Books = []string{"西游记", "水浒传", "红楼梦", "三国演义"}

func registerBookRoutes() {
	http.HandleFunc("/books", books)
	http.HandleFunc("/books/", book)
}

func books(w http.ResponseWriter, r *http.Request) {
	for i, b := range Books {
		fmt.Fprintln(w, i+1, b)

	}
}

func book(w http.ResponseWriter, r *http.Request) {
	// "r.URL.Path = /books/12"  12
	index, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/books/"))
	if err != nil || index >= len(Books) || index <= 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	fmt.Fprintln(w, Books[index-1])
}
