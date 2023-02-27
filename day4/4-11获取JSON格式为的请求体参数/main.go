package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe("127.0.0.1:8080", nil)
}

type User struct {
	User string `json:"user"`
	Pwd string `json:"pwd"`
}

func index(w http.ResponseWriter, r *http.Request) {
	user := User{}
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&user)
	fmt.Println(user)

	encoder := json.NewEncoder(w)
	encoder.Encode(user)
	
	 
}