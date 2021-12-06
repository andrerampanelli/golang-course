package main

import "fmt"

func main() {
	concept := scoreToConcept(7)
	fmt.Println(concept)
}

func scoreToConcept(score float64) string {
	// SWITCH is better
	if score >= 9 && score <= 10 {
		return "A"
	} else if score >= 8 && score < 9 {
		return "B"
	} else if score >= 5 && score < 8 {
		return "C"
	} else if score >= 3 && score < 5 {
		return "D"
	} else {
		return "E"
	}
}
