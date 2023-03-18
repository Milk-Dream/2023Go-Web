package main

import "database/sql"
import  _ "github.com/go-sql-driver/mysql"
import "fmt"

var db *sql.DB

func init() {
	dsn := "root:12345@tcp(127.0.0.1:3306)/my_db1"
	db, _ = sql.Open("mysql",dsn)
	fmt.Println(db)
	err := db.Ping()
	if err != nil {
		fmt.Println("链接失败", err)
	} else {
		fmt.Println("链接成功")
	}
}

func main() {
	//go get github.com/go-sql-driver/mysql
	fmt.Println(db.Stats())
}