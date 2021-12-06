package main

import "fmt"

func median(numbers ...float64) float64 {
	if len(numbers) == 0 {
		return 0.0
	}
	total := 0.0
	for _, num := range numbers {
		total += num
	}

	return total / float64(len(numbers))
}

func main() {
	fmt.Println(median(1.2, 2.3, 3.2, 5.4, 6.7))
}
