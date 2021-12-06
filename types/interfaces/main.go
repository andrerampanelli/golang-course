package main

import "fmt"

type Printable interface {
	toString() string
}

type Person struct {
	name     string
	lastName string
}

type Product struct {
	name  string
	price float64
}

func (p Person) toString() string {
	return fmt.Sprintf("%s %s", p.name, p.lastName)
}

func (p Product) toString() string {
	return fmt.Sprintf("%s - $ %.2f", p.name, p.price)
}

func print(x Printable) {
	fmt.Println(x.toString())
}

func main() {
	var thing Printable = Person{
		"Andre Vitor",
		"Rampanelli",
	}
	print(thing)

	thing = Product{
		"Pen",
		1.25,
	}
	print(thing)

	p2 := Person{
		"Arnold",
		"Schwarzenegger",
	}
	print(p2)

	print(Product{"Jeans Jacket", 54.99})
}
