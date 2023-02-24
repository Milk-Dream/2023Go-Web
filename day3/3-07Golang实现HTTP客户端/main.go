package main

import (
	
	"io/ioutil"
	"net/http"
)



func main() {
	resp, _ := http.Get("http://www.shicicn.com/zuozhe/129.html")
	body, _ := ioutil.ReadAll(resp.Body)
	//fmt.Println(string(body))
	ioutil.WriteFile("./index.html", body, 0666)


}