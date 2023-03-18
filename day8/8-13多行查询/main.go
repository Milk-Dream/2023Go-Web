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
	sqlStr := "select id, name, age from user where id>=?;"
	rows, _ := db.Query(sqlStr, 2)//这句话是把?替换成2
	for rows.Next() {
		var u User
		err := rows.Scan(&u.id, &u.name, &u.age)
		if err != nil {
			fmt.Println("查询失败", err)
			return
		}
		fmt.Println(u)
	}
	fmt.Println("查询结束")

}