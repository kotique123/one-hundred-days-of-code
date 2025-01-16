// Package phonenumber is a small library that implements function to work
// with phone numbers in North American format.
package phonenumber

import (
	"errors"
	"fmt"
)

// ErrInvalidPhoneNumber represents an error thrown in case of an invalid input.
var ErrInvalidPhoneNumber = errors.New("invalid phone number")

// Number returns cleanted up phoneNumber and an error value.
func Number(phoneNumber string) (string, error) {
	// Preallocates memory for a slice for a phone number
	number := make([]rune, 0, 11)
	for _, digit := range phoneNumber {
		if digit >= '0' && digit <= '9' {
			number = append(number, digit)
		}
		if len(number) > 10 {
			break
		}
	}
	// convrts slice back to string and assigns it back
	phoneNumber = string(number)
	// verifying format
	if len(phoneNumber) == 11 {
		if phoneNumber[0] == '1' {
			phoneNumber = phoneNumber[1:]
		} else {
			return "", ErrInvalidPhoneNumber
		}
	}
	if len(phoneNumber) != 10 {
		return "", ErrInvalidPhoneNumber
	} else if phoneNumber[0] == '0' || phoneNumber[0] == '1' {
		return "", ErrInvalidPhoneNumber
	} else if phoneNumber[3] == '0' || phoneNumber[3] == '1' {
		return "", ErrInvalidPhoneNumber
	}
	return phoneNumber, nil
}

// AreaCode returns a an area code from a phone number and an error value.
func AreaCode(phoneNumber string) (string, error) {
	phoneNumber, err := Number(phoneNumber)
	if err != nil {
		return "", err
	}
	// Returns just an area code.
	return phoneNumber[0:3], nil
}

// Format returns a phone number in specific format and an error value.
func Format(phoneNumber string) (string, error) {
	// Converts the phone number in format
	phoneNumber, err := Number(phoneNumber)
	if err != nil {
		return "", err
	}
	// Formats the number
	return fmt.Sprintf("(%s) %s-%s", phoneNumber[0:3], phoneNumber[3:6], phoneNumber[6:10]), nil
}
