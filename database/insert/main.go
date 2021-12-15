package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func panic(err error) {
	log.Fatal(err)
}

func main() {
	db, err := sql.Open("mysql", "root:root@/go_course")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	stmt, _ := db.Prepare("INSERT INTO users(name) VALUES (?)")
	stmt.Exec("Andre")
	stmt.Exec("Marina")

	res, _ := stmt.Exec("Hyago")
	id, _ := res.LastInsertId()
	rows, _ := res.RowsAffected()
	fmt.Println(id, rows)
}
