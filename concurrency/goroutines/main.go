package main

import (
	"fmt"
	"time"
)

func speak(person, text string, count int) {
	for i := 0; i < count; i++ {
		time.Sleep(time.Second)
		fmt.Printf("%s: %s (Iteration: %d)\n", person, text, i+1)
	}
}

func main() {
	go speak("Andre", "Speak with me!", 100)
	go speak("Monica", "You're boring!", 100)
}
