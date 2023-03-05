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

var temps *template.Template

func init() {
	temps = template.Must(template.ParseGlob("*.html"))
	/*tmp, err := template.ParseGlob("*.html")
	if err != nil {
		fmt.Println("模板解析失败")
		return
	}
	temps = tmp
	*/
}

func info(w http.ResponseWriter, r *http.Request) {
	tmp, _ := template.ParseFiles("base.html")
	tmp.Execute(w, nil)
	//temps.ExecuteTemplate(w, "base.html", nil)
}
func about(w http.ResponseWriter, r *http.Request) {
	//tmp, _ := template.ParseFiles("jd.html")
	//tmp.Execute(w, nil)
	//temps.ExecuteTemplate(w, "jd.html", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	//tmp, _ := template.ParseFiles("baidu.html")
	//tmp.Execute(w, nil)
	
	//顺序可以随便写
	//temps.ExecuteTemplate(w, "baidu.html", nil)
	//Execute的时候，顺序不能乱
	tmp, _ := template.ParseFiles("index.html")
	// tmp.Execute(w, "hello world")
	// tmp.Execute(w, "<h1>hello world</h1>")
	// tmp.Execute(w, "<script>alert('how are you?')</script>")
	tmp.Execute(w, "<script>alert('how are you?')</script>")

	//html js 可以安全转换为文本
	//直接将text/template改成html/template，所有的数据都会转换成字符串
}


func main() {
	http.HandleFunc("/test", info)
	http.HandleFunc("/index", index)
	// http.HandleFunc("/about", about)
	http.ListenAndServe("127.0.0.1:8080", nil)
}

