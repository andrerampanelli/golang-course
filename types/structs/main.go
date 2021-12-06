package main

import "fmt"

type Product struct {
	name        string
	description string
	price       float64
	discount    float64
}

// Receiver method
func (p Product) priceWithDiscount() float64 {
	return p.price * (1 - (p.discount / 100))
}

func main() {
	product := Product{
		"product",
		"description",
		159.99,
		10.0,
	}

	fmt.Println(product, product.priceWithDiscount())
}
