package main

import "fmt"

type Sport interface {
	turnTurboOn()
}

type Ferrari struct {
	name  string
	speed float64
	turbo bool
}

func (f *Ferrari) turnTurboOn() {
	f.turbo = true
}

func main() {
	car1 := Ferrari{"F40", 0.0, false}
	car1.turnTurboOn()

	var car2 Sport = &Ferrari{"458 Italia", 0.0, false} // Needs to be literal as type interface implemented
	car2.turnTurboOn()

	fmt.Println(car1, car2)
}
