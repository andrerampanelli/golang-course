package main

import (
	"fmt"
	"time"
)

// Generator
func speak(person string) <-chan string {
	ch := make(chan string)
	go func() {
		for i := 0; i < 3; i++ {
			ch <- fmt.Sprintf("%s speaking: %d\n", person, i)
			time.Sleep(time.Second)
		}
	}()
	return ch
}

// Multiplex - Mix string streaming into chan
func aggregate(in1, in2 <-chan string) <-chan string {
	ch := make(chan string)
	go func() {
		for {
			select {
			case s := <-in1:
				ch <- s
			case s := <-in2:
				ch <- s
			}
		}
	}()
	return ch
}

func main() {
	c := aggregate(speak("Andre"), speak("Monique"))

	fmt.Println(<-c, <-c)
	fmt.Println(<-c, <-c)
	fmt.Println(<-c, <-c)
}
