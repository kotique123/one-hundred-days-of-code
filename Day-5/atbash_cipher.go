// Package atbash is a small library which provides a function to encrypt a
// string using atbash cipher algorithm.
package atbash

import (
	"slices"
	"strings"
)

// Atbash encrypyts a string with an atbash cipher and returns a string.
func Atbash(s string) string {
	// convert string to lowercase
	s = strings.ToLower(s)

	// remove special characters
	runes := []rune(s)
	for i := 0; i < len(runes); i++ {
		if (runes[i] < 48 || runes[i] > 57) && (runes[i] < 97 || runes[i] > 122) {
			runes[i] = 32
		}
		// encrypt every character if the character is a letter and not a whitespace
		// for some reason whitespaces became encrypted as well.
		if !(runes[i] >= 48 && runes[i] <= 57) && runes[i] != 32 {
			runes[i] = 97 - runes[i] + 122
		}
	}
	// remove whitespaces
	no_whitespace := []rune("")
	for i := 0; i < len(runes); i++ {
		switch runes[i] {
		case 32:
			continue
		default:
			no_whitespace = append(no_whitespace, runes[i])
		}
	}
	// insert a whitespace after 5 characters
	for i := 5; i < len(no_whitespace); i += 6 {
		no_whitespace = slices.Insert(no_whitespace, i, 32)
	}
	return string(no_whitespace)
}

/*
Very cool solution made by bobahop (https://exercism.org/tracks/go/exercises/atbash-cipher/solutions/bobahop)
// Atbash converts the input string into Atbash cipher text.
func Atbash(input string) string {
	var output strings.Builder
	count := 0 // represents indexes
	for _, c := range []byte(input) {
		// checks lowercase letters and encrypts them
		if c > 96 && c < 123 {
			c = 'z' - (c - 'a')
		// checks uppercase letters and encrypts them
		} else if c > 64 && c < 91 {
			c = 'z' - ((c + 32) - 'a')
		// checks if the character is a number and ignores it
		} else if c < '0' || c > '9' {
			continue
		}
		// if index passed 5 characters, inserts a space to satisfy the requirements
		if count == 5 {
			output.WriteByte(' ')
			count = 0 // restore count to 0
		}
		output.WriteByte(c) // adds character to an output string
		count++ // index increment
	}
	return output.String() // returns an encrypted string
}
*/
