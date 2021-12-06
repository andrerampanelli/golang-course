package main

import "fmt"

func main() {
	first, second := switchNumber(2, 1)
	fmt.Printf("%d %d", first, second)
}

func switchNumber(p1, p2 int) (first int, second int) {
	first = p2
	second = p1
	return // Returns the named var declared in function scope
}
