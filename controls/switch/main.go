package main

import "time"

func main() {

}

func scoreToConcept(score float64) string {
	switch int(score) {
	case 10:
		// Other way!
		fallthrough
	case 9:
		return "A"
	case 8:
		return "B"
	case 7, 6, 5:
		return "C"
	case 4, 3:
		return "D"
	case 2, 1, 0:
		return "E"
	default:
		return "Invalid Score"
	}
}

func greetings() string {
	t := time.Now()
	switch { // Switch true
	case t.Hour() < 12:
		return "Good morning!"
	case t.Hour() < 18:
		return "Good afternoon!"
	default:
		return "Good evening!"
	}
}
