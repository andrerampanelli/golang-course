package main

import "fmt"

func main() {
	ch := make(chan int, 1)

	ch <- 1 // Write
	<-ch    // Read

	ch <- 2
	fmt.Println(<-ch)
}
