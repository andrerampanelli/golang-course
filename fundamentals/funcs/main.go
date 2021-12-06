package main

import (
	"fmt"
	"math"
)

func main() {
	print(sum(3, 4))
}

func sum(a int, b int) int {
	return a + b
}

func divide(a int, b int) int {
	return a / b
}

func subs(a int, b int) int {
	return a - b
}

func multiply(a int, b int) int {
	return a * b
}

func module(a int, b int) int {
	return a % b
}

func bitwiseAnd(a int, b int) int {
	return a & b
}

func bitwiseOr(a int, b int) int {
	return a | b
}

func bitwiseXor(a int, b int) int {
	return a ^ b
}

func min(a float64, b float64) float64 {
	return math.Min(a, b)
}

func max(a float64, b float64) float64 {
	return math.Max(a, b)
}

func power(a float64, b float64) float64 {
	return math.Pow(a, b)
}

func print(val int) {
	fmt.Println(val)
}
