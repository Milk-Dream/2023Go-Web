package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
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
var Publishers = []string {"人民出版社", "北京出版社"}


var temps = make(map[string]*template.Template, 5)
func init() {
	temps["home"] = template.Must(template.ParseFiles("./statics/htmls/home.html"))
	temps["list"] = template.Must(template.New("list.html").Funcs(template.FuncMap{"add": myAdd}).ParseFiles("./statics/htmls/home.html", "./statics/htmls/list.html"))
	temps["add"] = template.Must(template.ParseFiles("./statics/htmls/home.html", "./statics/htmls/add.html"))

	temps["edit"] = template.Must(template.New("edit.html").Funcs(template.FuncMap{"isIn":isIn}).ParseFiles("./statics/htmls/home.html", "./statics/htmls/edit.html"))

	temps["login"] = template.Must(template.ParseFiles("./statics/htmls/home.html", "./statics/htmls/login.html"))
}

func loginAuth(fn func(w http.ResponseWriter, r *http.Request)) http.HandlerFunc{
	return func (w http.ResponseWriter, r *http.Request)  {
		cookie, _ := r.Cookie("admin")
		if cookie == nil || cookie.Value != "adminvalue" {
			http.Redirect(w, r ,"/login", http.StatusTemporaryRedirect)
		}
		fn(w, r)
	}

}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/login", login)
	http.HandleFunc("/book_list", loginAuth(list))
	http.HandleFunc("/book_add",loginAuth(add))
	http.HandleFunc("/book_edit/",loginAuth(edit))
	http.HandleFunc("/book_delete/",loginAuth(delete))

	http.Handle("/statics/",http.FileServer(http.Dir(".")))
	http.ListenAndServe("127.0.0.1:8080", nil)
}

func myAdd(a, b int) int {
	return a + b
}

func isIn(a string ,list []string) bool {
	for _, s := range list {
		if s == a {
			return true
		}
	}
	return false
}

func login(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		tmp := temps["login"]
		tmp.Execute(w, nil)
	case "POST":
		name := r.PostFormValue("username")
		pawd := r.PostFormValue("password")
		if name == "admin" && pawd == "admin" {
			cookie := &http.Cookie{Name:"admin", Value: "adminvalue"}
			http.SetCookie(w, cookie)
			http.Redirect(w, r, "/book_list", http.StatusTemporaryRedirect)

		} else {
			tmp := temps["login"]
			tmp.Execute(w, "用户名或者密码错误")
		}

	}
}

func home(w http.ResponseWriter, r *http.Request) {
	tmp := temps["home"]
	tmp.Execute(w, nil)
}

func list(w http.ResponseWriter, r *http.Request) {
	//Books = append(Books, &Book{"人生", 19.8, "北京出版社", "2021-10-21", []string{"张三", "李四"}})
	 
	tmp := temps["list"]
	tmp.Execute(w, Books)
}

func add(w http.ResponseWriter, r *http.Request) {
	
	switch r.Method {
	case "GET":
		tmp := temps["add"]
		info := map[string][]string{"authors":Authors,"publisher":Publishers}

		tmp.Execute(w, info)
	case "POST":
		//保存数据，增加一个Book放在Books
		price, _ := strconv.ParseFloat(r.FormValue("price"),64)
		book := &Book{
			Title:r.FormValue("title"),
			Price: price,
			Publisher: r.FormValue("publisher"),
			PublishDate: r.FormValue("publish_date"),
			Authors: r.PostForm["authors"],
		}

		Books = append(Books, book)
		http.Redirect(w, r, "/book_list", 301)
	}
	
}

func edit(w http.ResponseWriter, r *http.Request) {

	id, _ := strconv.Atoi(r.URL.Path[len("/book_edit/"):])
	if id >= len(Books) {
		http.NotFound(w, r)
		return
	}
	switch r.Method {
	case "GET":
		
		tmp := temps["edit"]
		info := map[string]interface{}{"authors":Authors, "publishers": Publishers, "book": Books[id], "id": id}
		tmp.Execute(w, info)
	case "POST":
		//编辑索引
		book := Books[id]
		price, _ := strconv.ParseFloat(r.FormValue("price"), 64)
		book.Price = price
		book.Title = r.FormValue("title")
		book.PublishDate = r.FormValue("publish_date")
		book.Publisher = r.FormValue("publisher")
		book.Authors = r.PostForm["authors"]
		http.Redirect(w, r, "/book_list", http.StatusTemporaryRedirect)
	}
	
}

func delete(w http.ResponseWriter, r *http.Request) {
	fmt.Println("正在删除")
	id, _ := strconv.Atoi(r.URL.Path[len("/book_delete/"):])
	if id >= len(Books) {
		http.NotFound(w, r)
		return
	}
	//将索引为id的元素从Books删除
	Books = append(Books[:id],Books[id+1:]...)
	http.Redirect(w, r , "/book_list", http.StatusTemporaryRedirect)
}