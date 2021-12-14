package main

import (
	"fmt"
	"time"
)

func main() {

	base := 5
	ch := make(chan int, 1)

	go twoThreeFoutTimes(base, ch)

	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

func twoThreeFoutTimes(base int, ch chan int) {
	ch <- base * 2
	time.Sleep(time.Second)

	ch <- base * 3
	time.Sleep(time.Second)

	ch <- base * 4
}
