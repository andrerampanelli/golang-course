package main

import "math"

// PACKAGE only
// PUBLIC starts with first char uppercase
// PRIVATE starts with first char lowercase
// EXAMPLES:
// Dot is PUBLIC
// dot or _Dot is PRIVATE

// create dots in cartesian plan
type Dot struct {
	x float64
	y float64
}

func catheters(a, b Dot) (cx, cy float64) {
	cx = math.Abs(a.x - b.x)
	cy = math.Abs(a.y - b.y)
	return
}

// calculate the distance between catheters
func Distance(a, b Dot) float64 {
	cx, cy := catheters(a, b)
	return math.Sqrt(math.Pow(cx, 2) + math.Pow(cy, 2))
}
