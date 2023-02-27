package main

import (
	
	"io/ioutil"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe("127.0.0.1:8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	//上传文件
	//方式1:
	/*r.ParseMultipartForm(1024)
	fileHandler := r.MultipartForm.File["file"][0]
	file, _ := fileHandler.Open()
	body, _ := ioutil.ReadAll(file)
	fmt.Fprintln(w, body)
	ioutil.WriteFile("1.jpg", body, 0666)*/

	file, _, _ := r.FormFile("file")
	body, _ := ioutil.ReadAll(file)
	ioutil.WriteFile("1.jpg", body, 0666)
}