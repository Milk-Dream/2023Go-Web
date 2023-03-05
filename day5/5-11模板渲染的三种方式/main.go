package main

import (
	
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
	tmpArr, _ := template.ParseFiles("test.html", "jd.html", "baidu.html")

	/*
	tmpArr, _ = template.ParseGlob("*.html")
	// tmp.Execute(w, nil)
	tmpArr.ExecuteTemplate(w, "jd.html", nil)
	*/
	t := tmpArr.Lookup("jd.html")
	if t == nil {
		http.NotFound(w, r)
		return
	}
	t.Execute(w, nil)


}


func main() {
	http.HandleFunc("/test", info)
	
	http.ListenAndServe("127.0.0.1:8080", nil)
}

