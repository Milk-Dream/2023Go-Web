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
	
	//增加
	/*sqlStr := "insert into user(name, age) values(?,?)"
	ret, _ := db.Exec(sqlStr,"张三", 18)
	newId, _ := ret.LastInsertId()
	fmt.Println("新插入的数据的id是:",newId)
	*/

	//修改

	/*
	sqlStr := "update user set age=16 where id=1"
	ret, _ := db.Exec(sqlStr)
	counts, _ := ret.RowsAffected()
	fmt.Println(counts)
	*/
	//删除
	sqlStr := "delete from user where id=1"
	ret, _ := db.Exec(sqlStr)
	counts, _ := ret.RowsAffected()
	fmt.Println("影响的行数:",counts)

}