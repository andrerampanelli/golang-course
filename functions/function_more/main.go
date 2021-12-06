package main

import "fmt"

var sum = func(p1, p2 int) int {
	return p1 + p2
}

func main() {
	res := exec(pow, 3, 4)
	fmt.Println(res)
	res = sum(1, 2)
	fmt.Println(res)
}

func exec(function func(int, int) int, p1, p2 int) int {
	return function(p1, p2)
}

func pow(num1, num2 int) int {
	return num1 * num2
}
