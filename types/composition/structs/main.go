package main

import "fmt"

type Car struct {
	name     string
	topSpeed float64
}

type Ferrari struct {
	Car   // Anonymous field
	turbo bool
}

func main() {
	f := Ferrari{}
	f.name = "F40"
	f.topSpeed = 201.0
	f.turbo = true

	fmt.Printf("Ferrari %s have Top Speed equals %.2f mph and have turbo? %v", f.name, f.topSpeed, f.turbo)
}
