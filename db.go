package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type Product struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func main() {

	queryDb()

}

func queryDb() {
	fmt.Println("Go Mysql Query Test")
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/store")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	results, err := db.Query("Select id, name From products")

	if err != nil {
		panic(err.Error())
	}

	for results.Next() {
		var product Product

		err = results.Scan(&product.ID, &product.Name)

		if err != nil {
			panic(err.Error())
		}

		log.Printf(product.Name)
	}

}

func insertRow() {
	fmt.Println("Go Mysql Insert Test")
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
