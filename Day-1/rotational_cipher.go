package rotationalcipher

func RotationalCipher(plain string, shiftKey int) string {
	runes := []rune(plain)
	for i := 0; i < len(runes); i++ {
		if runes[i] >= 65 && runes[i] <= 90 {
			if runes[i]+rune(shiftKey) <= 90 {
				runes[i] = runes[i] + rune(shiftKey)
			} else {
				runes[i] = runes[i] + rune(shiftKey) - 26
			}
		} else if runes[i] >= 97 && runes[i] <= 122 {
			if runes[i]+rune(shiftKey) <= 122 {
				runes[i] = runes[i] + rune(shiftKey)
			} else {
				runes[i] = runes[i] + rune(shiftKey) - 26
			}
		}
	}
	return string(runes)
}
