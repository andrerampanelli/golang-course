package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	http.HandleFunc("/users/", UserHandler)

	log.Println("Listening ...")
	log.Fatal(http.ListenAndServe(":3000", nil))
}

type User struct {
	ID   int    `json:"ID"`
	Name string `json:"Name"`
}

func UserHandler(w http.ResponseWriter, r *http.Request) {
	sid := strings.TrimPrefix(r.URL.Path, "/users/")
	id, _ := strconv.Atoi(sid)

	switch {
	case r.Method == "GET" && id > 0:
		userByID(w, r, id)
	case r.Method == "GET":
		allUsers(w, r)
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Sorry... :(")
	}
}

func userByID(w http.ResponseWriter, r *http.Request, id int) {
	db, err := sql.Open("mysql", "root:root@/go_course")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	var user User
	db.QueryRow("SELECT id, name FROM users WHERE id = ?", id).Scan(&user.ID, &user.Name)

	json, _ := json.Marshal(user)

	w.Header().Set("Content-type", "application/json")
	fmt.Fprint(w, string(json))
}

func allUsers(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "root:root@/go_course")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, _ := db.Query("SELECT id, name FROM users")
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		rows.Scan(&user.ID, &user.Name)
		users = append(users, user)
	}

	json, _ := json.Marshal(users)
	w.Header().Set("Content-type", "application/json")
	fmt.Fprint(w, string(json))
}
