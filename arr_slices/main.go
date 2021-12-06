package main

import (
	"fmt"
)

func main() {
	array := [5]int{1, 2, 3, 4, 5}
	slice := []int{1, 2, 3, 4, 5}
	fmt.Print(array, slice)

	array[1] = 10
	slice = array[1:3]
	fmt.Print(array, slice)

	slice = append(slice, 4, 5)
	fmt.Print(array, slice)

	make_slice := make([]int, 5, 10)
	fmt.Print(make_slice, len(make_slice), cap(make_slice))

	make_slice = append(make_slice, 6, 7, 8, 9, 10, 11) // Auto increment based on started second argument passed in make function
	fmt.Print(make_slice, len(make_slice), cap(make_slice))
}
