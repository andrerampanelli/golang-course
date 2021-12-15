package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func panic(err error) {
	log.Fatal(err)
}

func update(db *sql.DB, name string, id int) {
	tx, _ := db.Begin()
	stmt, _ := tx.Prepare("UPDATE users SET name = ? WHERE id = ?")
	_, err := stmt.Exec(name, id)
	if err != nil {
		tx.Rollback()
		panic(err)
	}
	tx.Commit()
}

func delete(db *sql.DB, id int) {
	tx, _ := db.Begin()
	stmt, _ := tx.Prepare("DELETE FROM users WHERE id = ?")
	_, err := stmt.Exec(id)
	if err != nil {
		tx.Rollback()
		panic(err)
	}
	tx.Commit()
}

func main() {
	db, err := sql.Open("mysql", "root:root@/go_course")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	update(db, "Monique", 3)

	delete(db, 3)
}
