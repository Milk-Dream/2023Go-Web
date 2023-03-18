package main

import (
	
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)
type User struct {
	Id int 	`db:"id"`
	Pwd string `db:"name"`

	Name string `db:"pwd"`
}
var db *sqlx.DB

func init() {
	dsn := "root:12345@tcp(127.0.0.1:3306)/my_db1"
	db, _ = sqlx.Connect("mysql",dsn)

	fmt.Println(db)
	err := db.Ping()
	if err != nil {
		fmt.Println("链接失败", err)
	} else {
		fmt.Println("链接成功")
	}
}
func main() {
	//go get github.com/jmoiron/sqlx
	var u User
	sqlStr := "select id, name, pwd from user where id=?"
	/*
	err := db.QueryRow(sqlStr,1).Scan(&u.id,&u.name,&u.pwd)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(u)
	*/
	err := db.Get(&u, sqlStr, 1)
	if(err != nil) {
		fmt.Println(err)
		return
	}
	fmt.Println(u)

	var users []User
	err = db.Select(&users, sqlStr, 1)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(users)

	
}