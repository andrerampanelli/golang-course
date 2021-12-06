package main

import "fmt"

func main() {
	people := make(map[int]string)

	people[123456789] = "Person #1"
	people[213546879] = "Person #2"

	for cod, name := range people {
		fmt.Printf("%s (Cod: %d)\n", name, cod)
	}

	other_people := map[int]string{
		1: "Person #1",
		2: "Person #2",
	}

	for _, name := range other_people {
		fmt.Printf("%s\n", name)
	}
}
