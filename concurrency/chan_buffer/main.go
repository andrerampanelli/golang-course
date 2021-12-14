package main

import "fmt"

func routine(ch chan int) {
	ch <- 1
	ch <- 2
	ch <- 3
	// Blocks, buffer full
	ch <- 4
	fmt.Println("Running ...") // Not running because of buffer
	ch <- 5
	ch <- 6
}

func main() {
	ch := make(chan int, 3)
	go routine(ch)

	fmt.Println(<-ch)
}
