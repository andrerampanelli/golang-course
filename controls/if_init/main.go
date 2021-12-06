package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	if i := randInt(); i > 5 { // Switch support
		fmt.Println("You won!")
		return
	}
	fmt.Println("You lose!")
}

func randInt() int {
	src := rand.NewSource(time.Now().UnixNano())
	randInt := rand.New(src)
	return randInt.Intn(10)
}
