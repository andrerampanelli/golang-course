package main

import (
	"fmt"
	"math"
)

func main() {
	const PI float64 = 3.1415926537
	var raio = 3.2 // Compiler auto detecting type float64

	// Var not init is init by ":="
	area := PI * math.Pow(raio, 2)

	fmt.Println("area is", area)

	// Multiple creating
	const (
		a = 1
		b = 2
	)

	var (
		c = 3
		d = 4
	)
	fmt.Println(a, b, c, d)

	var e, f bool = true, false
	fmt.Println(e, f)

	g, h, i := 2, false, "hey!"
	fmt.Println(g, h, i)
}
