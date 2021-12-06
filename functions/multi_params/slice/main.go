package main

import "fmt"

func printApproveds(approveds ...string) {
	fmt.Println("Approved List")
	for _, approved := range approveds {
		fmt.Println(approved)
	}
}

func main() {
	approveds := []string{"Elba", "Raul", "Michael", "Alexander "}
	printApproveds(approveds...)
}
