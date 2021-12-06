package main

import "fmt"

func main() {
	printApprovedStatus(6)
}

func printApprovedStatus(score int) {
	if score >= 7 {
		fmt.Println("Approved with score:", score)
	} else {
		fmt.Println("Not approved with score:", score)
	}
}
