package main

import (
	"net/http"
	"strconv"
	"text/template"
)

type Book struct {
	Title string
	Price float64
	Publisher string
	PublishDate string
	Authors []string
}

var Books []*Book
var Authors = []string{"张三", "李四"}
var Publishers = []string{"人民出版社", "北京出版社"}


func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/book_add", add)
	http.HandleFunc("/book_edit", edit)
	http.HandleFunc("/book_delete", delete)
	http.HandleFunc("/book_list", list)
	http.Handle("/statics/", http.FileServer(http.Dir(".")))
	http.ListenAndServe("localhost:8888", nil)
}

func list(w http.ResponseWriter, r *http.Request) {
	Books = append(Books, &Book{"人生", 19.8, "北京出版社", "2021-10-22", []string{"张三", "李四"}})
	//tmp, _ := template.ParseFiles("home.html", "list.html")
	tmp, _ := template.New("list.html").Funcs(template.FuncMap{"add":addOne}).ParseFiles("home.html", "list.html")
	tmp.Execute(w,Books)
	
}

func addOne(a, b int) int{
	return a + b
}

func isIn(a string, list []string) bool {
	for _, s := range list {
		if s == a {
			return true
		}
	}

	return false

}

func add(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		tmp, _ := template.ParseFiles("home.html", "add.html")
		info := map[string][]string {"authors":Authors,"publishers":Publishers}
		tmp.ExecuteTemplate(w ,"add.html", info)
	case "POST":
		//保存数据
		price, _ := strconv.ParseFloat(r.FormValue("price"), 64)
		book := &Book {
			Title: r.FormValue("title"),
			Price: price,
			Publisher: r.FormValue("publisher"),
			PublishDate: r.FormValue("publish_date"),
			Authors: r.PostForm["authors"],
		}
		Books = append(Books, book)
		http.Redirect(w, r, "/list", 301)

	}
	
	
}

func edit(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Path[len("book_edit/"):]
	id, _ := strconv.Atoi(idStr)
	if(id >= len(Books)) {
		http.NotFound(w, r)
		return
	}
	switch r.Method {
	case "GET":
		tmp, _ := template.New("edit.html").Funcs(template.FuncMap{"isIn":isIn}).ParseFiles("home.html", "edit.html")
		info := map[string]interface{} {"authors":Authors,"publishers":Publishers,"id":id,"books":Books[id]}
		tmp.ExecuteTemplate(w ,"edit.html", info)
	case "POST":
		//保存数据
		book := Books[id]
		price, _ := strconv.ParseFloat(r.FormValue("price"), 64)

		book.Price = price
		book.Title = r.FormValue("title")
		book.PublishDate = r.FormValue("publish_date")
		book.Publisher = r.FormValue("publisher")
		book.Authors = r.PostForm["authors"]
		http.Redirect(w, r, "/list", 301)

	}
	
	
}

func delete(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Path[len("book_edit/"):]
	id, _ := strconv.Atoi(idStr)
	if(id >= len(Books)) {
		http.NotFound(w, r)
		return
	}
	Books = append(Books[:id],Books[id+1:]...)
	http.Redirect(w, r, "/book_list", 301)
}

func home(w http.ResponseWriter, r *http.Request) {
	tmp, _ := template.ParseFiles("home.html")
	tmp.Execute(w, nil)
}