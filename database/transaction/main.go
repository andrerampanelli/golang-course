package main

import (
	"database/sql"
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

	tx, _ := db.Begin()
	stmt, _ := tx.Prepare("INSERT INTO users(id, name) VALUES (?, ?)")
	stmt.Exec(200, "Andre")
	stmt.Exec(201, "Marina")

	_, err = stmt.Exec(1, "Hyago") // Duplicate ID
	if err != nil {
		tx.Rollback()
		panic(err)
	}

	tx.Commit()
}
