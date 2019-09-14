package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println("Go Mysql Test")

	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/store")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	insert, err := db.Query("INSERT INTO products(name, price) VALUES ('Product 1', '10')")

	if err != nil {
		panic(err.Error())
	}

	defer insert.Close()

}
