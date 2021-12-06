package main

import (
	"encoding/json"
	"fmt"
)

type Product struct {
	ID    string   `json:"id"`
	Name  string   `json:"name"`
	Price float64  `json:"price"`
	Tags  []string `json:"tags"`
}

func main() {
	// Struct to JSON
	p1 := Product{
		"df14b7ce-1a4e-48ea-a034-b722e11765a4",
		"Macbook Pro M1",
		1000.00,
		[]string{"Notebook", "Apple", "Macbook", "M1"},
	}

	p1JSON, _ := json.Marshal(p1)
	fmt.Println(string(p1JSON))

	var p2 Product
	jsonString := `{"id":"fe6d0bb7-5363-4067-991c-c18fbb7eaba4","name":"Pencil","price":1.25,"tags":["Stationery","Writing"]}`
	json.Unmarshal([]byte(jsonString), &p2)

	fmt.Println(p2)
}
