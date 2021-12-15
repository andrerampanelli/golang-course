package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func currentTime(w http.ResponseWriter, r *http.Request) {
	s := time.Now().Format("02/01/2006 03:04:05")
	fmt.Fprintf(w, "<h1>Current time: %s</h1>", s)
}

func main() {
	http.HandleFunc("/current-time", currentTime)

	log.Println("Listening . . .")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
