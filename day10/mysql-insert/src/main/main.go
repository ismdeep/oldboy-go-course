package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var Db *sqlx.DB

func init() {
	database, err := sqlx.Open("mysql", "root:root@tcp(127.0.0.1)/testdb")
	if err != nil {
		fmt.Println("Open MySQL failed, err:", err)
		return
	}
	Db = database
}

func main() {
	r, err := Db.Exec("insert into person(username, sex, email) values(?,?,?)", "stu001", "man", "stu01@qq.com")
	if err != nil {
		fmt.Println("Exec failed, err:", err)
	}
	id, err := r.LastInsertId()
	if err != nil {
		fmt.Println("Exec failed, err:", err)
		return
	}
	fmt.Println("insert successfully. Id:", id)
}
