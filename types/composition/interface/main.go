package main

import "fmt"

type Sport interface {
	turnTurboOn()
}

type Luxuous interface {
	autoParking()
}

// Composition
type LuxuousSport interface {
	Sport
	Luxuous
	// And other if needed
}

type BMW_M7 struct{}

func (b BMW_M7) turnTurboOn() {
	fmt.Println("VU TUTUTUTU")
}

func (b BMW_M7) autoParking() {
	fmt.Println("Parking ...")
}

func main() {
	var b LuxuousSport = BMW_M7{}

	b.turnTurboOn()
	b.autoParking()
}
