package main

import "fmt"

type Course struct {
	name string
}

func main() {
	var thing interface{}
	fmt.Println(thing)

	thing = 3
	fmt.Println(thing)

	type Dynamic interface{}
	var thing2 Dynamic = "Hello"
	fmt.Println(thing2)

	thing2 = true
	fmt.Println(thing2)

	thing2 = Course{"Golang is aweasome"}
	fmt.Println(thing2)
}
