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
	/*
		CREATE TABLE `user` (
			`id` BIGINT(20) NOT NULL AUTO_INCREMENT,
			`name` VARCHAR(20) DEFAULT '',
			`balance` INT(11),
			PRIMARY KEY(`id`)
		)ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;
	*/

	tx , _:= db.Begin()
	sql1 := "update user set balance=400 where name='张三'"

	ret, err := tx.Exec(sql1)
	if ret != nil {
		fmt.Println("张三账户转账失败", err)
		tx.Rollback()
		return

	}
	fmt.Println("张三账户减去了100元，账户剩余400元")

	
	//李四账户增加100元
	sql2 := "update user set balance=500 where name='李三'"
	ret2, err := tx.Exec(sql2)
	if ret2 != nil {
		fmt.Println("李四账户余额增加失败了")
		tx.Rollback()
		return
	}

	fmt.Println("李四账户加了100元，账户剩余500元")
	tx.Commit()

}