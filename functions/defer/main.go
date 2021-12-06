package main

import "fmt"

func main() {
	fmt.Println(getApprovedSpeed(6000))
}

func getApprovedSpeed(num int) int {
	defer fmt.Println("End")
	if num > 5000 {
		fmt.Println("You're going too fast")
		return 5000
	}
	fmt.Println("You're too slow")
	return num
}
