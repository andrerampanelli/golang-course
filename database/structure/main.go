package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func panic(err error) {
	log.Fatal(err)
}

func exec(db *sql.DB, sql string) sql.Result {
	result, err := db.Exec(sql)
	if err != nil {
		panic(err)
	}
	return result
}

func main() {
	db, err := sql.Open("mysql", "root:root@/")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	exec(db, "CREATE DATABASE IF NOT EXISTS go_course")
	exec(db, "USE go_course")
	exec(db, "DROP TABLE IF EXISTS users")
	exec(db, `CREATE TABLE users (
		id integer auto_increment,
		name varchar(80),
		PRIMARY KEY(id)
	)`)
}
