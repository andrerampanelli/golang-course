package main

import "fmt"

type Item struct {
	_id      string
	quantity int
	price    float64
}

type Order struct {
	_id   string
	items []Item
}

func (o Order) finalPrice() float64 {
	total := 0.0

	for _, item := range o.items {
		total += item.price * float64(item.quantity)
	}

	return total
}

func main() {
	order := Order{
		_id: "9482b7de-b561-44f4-98bb-66e75132dde6",
		items: []Item{
			{
				_id:      "3f89fe15-4b6d-440d-bda0-d2062894772f",
				quantity: 5,
				price:    10.0,
			},
			{
				_id:      "39ae0cbf-2f12-44e6-8b2c-6caf548c26c0",
				quantity: 2,
				price:    5.5,
			},
			{
				_id:      "df14b7ce-1a4e-48ea-a034-b722e11765a4",
				quantity: 1,
				price:    1345.37,
			},
		},
	}

	fmt.Printf("Final price is: $ %.2f", order.finalPrice())
}
