package main

import "fmt"

func main() {
	fact, _ := factorial(5)
	fmt.Println(fact)

	_, err := factorial(-4)
	if err != nil {
		fmt.Println(err)
	}

	res := betterFactorial(5)
	fmt.Println(res)
}

func factorial(n int) (int, error) {
	switch {
	case n < 0:
		return -1, fmt.Errorf("Invalid number: %d", n)
	case n == 0:
		return 1, nil
	default:
		fact, _ := factorial(n - 1)
		return n * fact, nil
	}
}

// Thats a better way
func betterFactorial(n uint) uint {
	switch {
	case n == 0:
		return 1
	default:
		return n * betterFactorial(n-1)
	}
}
