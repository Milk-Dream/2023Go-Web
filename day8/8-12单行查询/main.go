package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

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

type User struct {
	id int
	age int
	name string
}

func main() {
	//go get github.com/go-sql-driver/mysql
	fmt.Println(db.Stats())
	/*
		CREATE DATABASE my_db1;
use my_db1;
CREATE TABLE `user` (
    `id` BIGINT(20) NOT NULL AUTO_INCREMENT,
    `name` VARCHAR(20) DEFAULT '',
    `age` INT(11) DEFAULT '0',
    PRIMARY KEY(`id`)
)ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;
	*/
	var u User
	sqlStr := "select id, name, age from user where id=3"
	err := db.QueryRow(sqlStr).Scan(&u.id, &u.name, &u.age)
	if err != nil {
		fmt.Println("查询失败:",err)
	}
	fmt.Println(u)

}