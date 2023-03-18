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
	pwd string

	name string
}

func main() {
	/*
	CREATE TABLE `user` (
		`id` BIGINT(20) NOT NULL AUTO_INCREMENT,
		`name` VARCHAR(20) DEFAULT '',
		`pwd` VARCHAR(20) DEFAULT '',
		PRIMARY KEY(`id`)
	)ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;
	*/
	name := "jack"
	pwd := "123456"
	sqlStr := "select id, name, pwd from user where name=? and pwd=?"
	//sql注入 name = 'xxx' or 1=1#
	sqlStr1 := fmt.Sprintf("select id, name, pwd from user where name='%s' and pwd='%s'",name,pwd)
	fmt.Println(sqlStr1)
	//安全的做法
	rows, _ := db.Query(sqlStr,name,pwd)
	for rows.Next() {
		var u User
		err := rows.Scan(&u.id, &u.name, &u.pwd)
		if err != nil {
			fmt.Println("登录失败",err)
			return
		}
		fmt.Println(u)
	}

}