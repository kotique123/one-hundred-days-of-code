// Package rotationalcipher is a small library that can encrypt a string by a shift in English alphabet
package rotationalcipher

// RotationalCipher returns a string which is encrypted by a shiftKey
func RotationalCipher(plain string, shiftKey int) string {
	// convert string to a slice of runes
	runes := []rune(plain)
	// iterate through string
	for i := 0; i < len(runes); i++ {
		// check if a rune is an uppercase latin letter based on utf-8 decimal value
		if runes[i] >= 65 && runes[i] <= 90 {
			// determine if shift is out of bounds of alphabet
			if runes[i]+rune(shiftKey) <= 90 {
				runes[i] = runes[i] + rune(shiftKey)
			} else {
				runes[i] = runes[i] + rune(shiftKey) - 26
			}
		// check if a rune is a lowercase latin letter based on utf-8 decimal value
		} else if runes[i] >= 97 && runes[i] <= 122 {
			// determine if shift is out of bounds of alphabe
			if runes[i]+rune(shiftKey) <= 122 {
				runes[i] = runes[i] + rune(shiftKey)
			} else {
				runes[i] = runes[i] + rune(shiftKey) - 26
			}
		}
	}
	// convert a slice of runes back to string and return it
	return string(runes)
}
