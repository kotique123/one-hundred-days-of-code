// Package protein is a small library that implements function to translate proteins.
package protein

import (
	"errors"
)

// ErrStop is a variable that represents an error when encounters a stop codon.
var ErrStop = errors.New("stop sequence")

// ErrInvalidBase is a variable that represents an error when encounters invalid RNA sequence.
var ErrInvalidBase = errors.New("invalid protein base")

// FromRNA takes an RNA sequence and returns a slice of proteins and an error value.
func FromRNA(rna string) ([]string, error) {
	if len(rna)%3 != 0 {
		return nil, ErrInvalidBase
	} else {
		var proteins []string
		start := 0
		end := 3
		var status error
		var acid string
		for end <= len(rna) {
			acid, status = FromCodon(rna[start:end])
			if status == nil {
				proteins = append(proteins, acid)
				start += 3
				end += 3
			} else {
				return proteins, nil
			}
		}
		return proteins, nil
	}

}

// FromCodon takes a codon string and returns name of protein and an error value.
func FromCodon(codon string) (string, error) {
	switch codon {
	case "AUG":
		return "Methionine", nil
	case "UUU", "UUC":
		return "Phenylalanine", nil
	case "UUA", "UUG":
		return "Leucine", nil
	case "UCU", "UCC", "UCA", "UCG":
		return "Serine", nil
	case "UAU", "UAC":
		return "Tyrosine", nil
	case "UGU", "UGC":
		return "Cysteine", nil
	case "UGG":
		return "Tryptophan", nil
	case "UAA", "UAG", "UGA":
		return "", ErrStop
	default:
		return "", ErrInvalidBase
	}
}
