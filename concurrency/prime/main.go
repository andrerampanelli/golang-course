package main

import (
	"fmt"
	"time"
)

func isPrime(num int) bool {
	for i := 2; i < num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func primes(count int, ch chan int) {
	start := 2
	for i := 0; i < count; i++ {
		for prime := start; ; prime++ {
			if isPrime(prime) {
				ch <- prime
				start = prime + 1
				time.Sleep(time.Second)
				break
			}
		}
	}
	close(ch)
}

func main() {
	ch := make(chan int, 30)

	go primes(cap(ch), ch)

	for primes := range ch {
		fmt.Printf("%d\n", primes)
	}
}
