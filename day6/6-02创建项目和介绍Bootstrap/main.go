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

var temps = make(map[string]*template.Template, 8)

//全局
func loginAuth(fn func(w http.ResponseWriter, r *http.Request)) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		cookie, _ := r.Cookie("admin")
		if cookie == nil || cookie.Value != "adminValue" {
			http.Redirect(w, r, "/login", 301)
		}
		fn(w, r)
	}

}

func init() {

	addOne := func (a, b int) int{
		return a + b
	}

	isIn := func (a string, list []string) bool {
		for _, s := range list {
			if s == a {
				return true
			}
		}
	
		return false
	
	}
	temps["home"] = template.Must(template.ParseFiles("./statics/htmls/home.html"))
	temps["list"] = template.Must(template.New("./statics/htmls/list.html").Funcs(template.FuncMap{"add":addOne}).ParseFiles("./statics/htmls/home.html", "./statics/htmls/list.html"))

	temps["edit"] = template.Must(template.New("./statics/htmls/edit.html").Funcs(template.FuncMap{"isIn":isIn}).ParseFiles("./statics/htmls/home.html", "./statics/htmls/edit.html"))
	temps["login"] = template.Must(template.ParseFiles("./statics/htmls/login.html"))
}

func login(w http.ResponseWriter, r *http.Request) {
	tmp := temps["login"]
	switch r.Method {
	case "GET":
		
		tmp.Execute(w, nil)
	case "POST":
		//登录，校验选项
		name := r.PostFormValue("username")
		pawd := r.PostFormValue("password")
		if name == "admin" && pawd == "admin" {
			//设置cookie
			cookie := &http.Cookie{Name: "admin", Value: "adminValue"}
			http.SetCookie(w, cookie)
			http.Redirect(w, r, "/book_list", 301)
		} else {
			
			tmp.Execute(w, "用户名和密码错误")
		}

	}
}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/book_add", loginAuth(add))
	http.HandleFunc("/book_edit", loginAuth(edit))
	http.HandleFunc("/book_delete", loginAuth(delete))
	http.HandleFunc("/book_list", list)
	http.HandleFunc("/login", login)
	http.Handle("/statics/", http.FileServer(http.Dir(".")))
	http.ListenAndServe("localhost:8888", nil)
}

func list(w http.ResponseWriter, r *http.Request) {
	Books = append(Books, &Book{"人生", 19.8, "北京出版社", "2021-10-22", []string{"张三", "李四"}})
	//tmp, _ := template.ParseFiles("home.html", "list.html")
	tmp, _ := temps["list"]
	tmp.Execute(w,Books)
	
}

func add(w http.ResponseWriter, r *http.Request) {
	//校验cookie
	cookie, _ := r.Cookie("admin")
	if cookie == nil || cookie.Value != "adminValue" {
		http.NotFound(w, r)
		return
	}
	
	switch r.Method {
	case "GET":
		tmp := temps["add"]
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
		tmp := temps["edit"]
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
	// tmp, _ := template.ParseFiles("home.html")
	tmp := temps["home"]
	tmp.Execute(w, nil)
}