package main

import "fmt"

func main() {
	first()
	second("second")
	// third := third()
	// fourth := fourth("1", "2")

	third, fourth := third(), fourth("1", "2")

	fmt.Printf("3) %s\n 4) %s", third, fourth)
}

func first() {
	fmt.Println("First")
}

func second(param1 string) {
	fmt.Printf("Second: %s\n", param1)
}

func third() string {
	return "Third"
}

func fourth(param1, param2 string) string {
	return fmt.Sprintf("Fourth:\n 1) %s\n 2) %s", param1, param2)
}
