package main

import (
	"io/ioutil"
	"net/http"
	"text/template"
)

type User struct {
	Name string
	Age int
}

type Student struct {
	Name string
	Age int
	Scores map[string]int
}

func info(w http.ResponseWriter, r *http.Request) {
	tmp := template.New("test.html")
	b, _ := ioutil.ReadFile("test.html")
	tmp.Parse(string(b))
	tmp.Execute(w, nil)

	//ParseFiles
	tmp, _ = template.ParseFiles("test.html")
	tmp.Execute(w, nil)

	//ParseGlob，以通配符
	tmp, _ = template.ParseGlob("*.html")
	tmp.Execute(w, nil)

}


func main() {
	http.HandleFunc("/test", info)
	
	http.ListenAndServe("127.0.0.1:8080", nil)
}

