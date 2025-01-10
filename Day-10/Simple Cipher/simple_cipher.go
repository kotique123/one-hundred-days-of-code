// Package cipher is a small library that allows to encrypt strings based on
// a shift or a key.

package cipher

import (
	"slices"
)

type shift int       // shift represent a shift in a cipher.
type vigenere string // vigenere represents a key for a cipher.

// NewCesar returns a Cipher with a default shift of three accroding to Cesar's cipher.
func NewCaesar() Cipher {
	return NewShift(3)
}

// NewShift returns a Cipher type based on distance if it is valid.
func NewShift(distance int) Cipher {
	if distance > 25 || distance == 0 || distance < -25 {
		return nil
	}
	if distance < 0 {
		distance += 26
	}
	return shift(distance)
}

// Encode returns an encoded string based on shift.
func (c shift) Encode(input string) string {
	runes := []rune(input)
	for i := 0; i < len(runes); i++ {
		if runes[i] >= 65 && runes[i] <= 90 {
			runes[i] = runes[i] + 32
			shift_index := int(runes[i]) + int(c)
			if shift_index > 122 {
				shift_index -= 26
			}
			runes[i] = rune(shift_index)
		} else if runes[i] >= 96 && runes[i] <= 122 {
			shift_index := int(runes[i]) + int(c)
			if shift_index > 122 {
				shift_index -= 26
			}
			runes[i] = rune(shift_index)
		} else {
			runes = slices.Delete(runes, i, i+1)
			i--
		}
	}
	return string(runes)
}

// Decode returns a decoded string based on shift.
func (c shift) Decode(input string) string {
	runes := []rune(input)
	for i := 0; i < len(runes); i++ {
		if runes[i] >= 65 && runes[i] <= 90 {
			runes[i] = runes[i] + 32
			shift_index := int(runes[i]) - int(c)
			if shift_index < 97 {
				shift_index += 26
			}
			runes[i] = rune(shift_index)
		} else if runes[i] >= 96 && runes[i] <= 122 {
			shift_index := int(runes[i]) - int(c)
			if shift_index < 97 {
				shift_index += 26
			}
			runes[i] = rune(shift_index)
		} else {
			return ""
		}
	}
	return string(runes)
}

// NewVigenere returns a Cipher type of a key if it is valid.
func NewVigenere(key string) Cipher {
	matches_format := true
	not_all_a := false
	for i := 0; i < len(key); i++ {
		if key[i] != 97 {
			not_all_a = true
		}
		if key[i] >= 97 && key[i] <= 122 {
			matches_format = true
		} else {
			matches_format = false
			break
		}
	}
	if matches_format && not_all_a {
		return vigenere(key)
	} else {
		return nil
	}
}

// Encode encodes an input string based on a key and returns a string.
func (v vigenere) Encode(input string) string {
	runes_input := []rune(input)
	runes_vigenere := []rune(v)

	vigeneres := make(map[int]int, len(runes_vigenere))

	for i := 0; i < len(runes_vigenere); i++ {
		vigeneres[i] = int(runes_vigenere[i]) - 97
	}

	map_index := 0

	for i := 0; i < len(runes_input); i++ {

		if map_index == len(vigeneres) {
			map_index = 0
		}
		if runes_input[i] >= 65 && runes_input[i] <= 90 {
			runes_input[i] += 32
			shift_index := int(runes_input[i]) + int(vigeneres[map_index])
			if shift_index > 122 {
				shift_index -= 26
			}
			runes_input[i] = rune(shift_index)
			map_index++

		} else if runes_input[i] >= 96 && runes_input[i] <= 122 {
			shift_index := int(runes_input[i]) + int(vigeneres[map_index])
			if shift_index > 122 {
				shift_index -= 26
			}
			runes_input[i] = rune(shift_index)
			map_index++

		} else {
			runes_input = slices.Delete(runes_input, i, i+1)
			i--

		}
	}
	return string(runes_input)
}

// Decode returns an encoded string based on a key.
func (v vigenere) Decode(input string) string {
	runes_input := []rune(input)
	runes_vigenere := []rune(v)

	vigeneres := make(map[int]int, len(runes_vigenere))

	for i := 0; i < len(runes_vigenere); i++ {
		vigeneres[i] = int(runes_vigenere[i]) - 97
	}

	map_index := 0

	for i := 0; i < len(runes_input); i++ {

		if map_index == len(vigeneres) {
			map_index = 0
		}
		if runes_input[i] >= 65 && runes_input[i] <= 90 {
			runes_input[i] += 32
			shift_index := int(runes_input[i]) - int(vigeneres[map_index])
			if shift_index < 97 {
				shift_index += 26
			}
			runes_input[i] = rune(shift_index)
			map_index++

		} else if runes_input[i] >= 96 && runes_input[i] <= 122 {
			shift_index := int(runes_input[i]) - int(vigeneres[map_index])
			if shift_index < 97 {
				shift_index += 26
			}
			runes_input[i] = rune(shift_index)
			map_index++

		} else {
			runes_input = slices.Delete(runes_input, i, i+1)
			i--

		}
	}
	return string(runes_input)
}
