package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type user struct {
	id   int
	name string
}

func panic(err error) {
	log.Fatal(err)
}

func main() {
	db, err := sql.Open("mysql", "root:root@/go_course")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, _ := db.Query("SELECT id, name FROM users WHERE id > ?", 0)
	defer rows.Close()

	for rows.Next() {
		var u user
		rows.Scan(&u.id, &u.name)
		fmt.Println(u)
	}
}
