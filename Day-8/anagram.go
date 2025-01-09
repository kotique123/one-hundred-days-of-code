// Package angram is a small library that has functions to find angrams of a word.
package anagram

import "strings"

// CreateMapOfSubjects returns a map of integers with number of the occurencies
// in the subject.
func CreateMapOfSubject(subject string) map[int]int {
	subject_map := make(map[int]int)
	for i := 0; i < len(subject); i++ {
		subject_map[int(subject[i])]++
	}
	return subject_map
}

// CompareTwoMaps compares two maps based on a given string.
func CompareTwoMaps(word string, subject, candidate map[int]int) bool {
	for i := 0; i < len(word); i++ {
		if subject[int(word[i])] == candidate[int(word[i])] {
			continue
		} else {
			return false
		}
	}
	return true
}

// Detect returns a slice of words that are angrams based on the subject
// and a slice of candidates.
func Detect(subject string, candidates []string) []string {
	subject_map := CreateMapOfSubject(strings.ToLower(subject))
	var anagrams []string
	for _, word := range candidates {
		if len(word) == len(subject) && !strings.EqualFold(word, subject) {
			subject_map2 := CreateMapOfSubject(strings.ToLower(word))
			if CompareTwoMaps(strings.ToLower(subject), subject_map, subject_map2) {
				anagrams = append(anagrams, word)
			}
		}

	}
	return anagrams
}
