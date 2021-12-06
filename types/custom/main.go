package main

import "fmt"

type Score float64

func (n Score) between(start, end float64) bool {
	return float64(n) >= start && float64(n) <= end
}

func scoreToConcept(n Score) string {
	if n.between(9.0, 10.0) {
		return "A"
	} else if n.between(7.0, 8.99) {
		return "B"
	} else if n.between(5.0, 7.99) {
		return "C"
	} else if n.between(3.0, 4.99) {
		return "D"
	} else {
		return "E"
	}
}

func main() {
	fmt.Println(scoreToConcept(2.1))
	fmt.Println(scoreToConcept(6.3))
	fmt.Println(scoreToConcept(9.5))
}
