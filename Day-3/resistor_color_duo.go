// Package resistorcolorduo is a small library that has a function to determine
// resistor value based on color coding.
package resistorcolorduo

// Value returns a value of resistor based on color coding.
func Value(colors []string) int {
	result := 0
	multipler := 10
	for i := 0; i < 2; i++ {

		switch colors[i] {
		case "black":
			result += 0 * multipler
		case "brown":
			result += 1 * multipler
		case "red":
			result += 2 * multipler
		case "orange":
			result += 3 * multipler
		case "yellow":
			result += 4 * multipler
		case "green":
			result += 5 * multipler
		case "blue":
			result += 6 * multipler
		case "violet":
			result += 7 * multipler
		case "grey":
			result += 8 * multipler
		case "white":
			result += 9 * multipler
		}
		multipler -= 9
	}
	return result
}
