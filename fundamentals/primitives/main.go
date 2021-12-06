package main

func main() {
	// Numbers
	// Whole Numbers (W) := uint8(byte ), uint16, uint32, uint64 => 0
	// Integers (Z) := int8, int16, int32 (rune *unicode ), int64 => 0

	// Real numbers (R) := float32, float64 => 0

	// Booleans := bool => false

	// Strings := string => "", `` => (len()) => ""

	// Pointers *(type) => nil
	// pointer = &var
	// *pointer => val

	// Converts

	// str -> int
	// str, err := strconv.Atoi(num)
	// int -> str
	// num, err := strconv.Itoa(str)
	// parse bool ("true" or "1") => true | false
	// boo, err = strconv.ParseBool(str)
}
