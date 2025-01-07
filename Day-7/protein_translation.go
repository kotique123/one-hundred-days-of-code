package protein

import (
	"errors"
)

var ErrStop = errors.New("stop sequence")
var ErrInvalidBase = errors.New("invalid protein base")

func FromRNA(rna string) ([]string, error) {
	if len(rna)%3 != 0 {
		return nil, ErrInvalidBase
	} else {
		var proteins []string
		start := 0
		end := 3
		index := 0
		var status error
		var acid string
		for end <= len(rna) {
			acid, status = FromCodon(rna[start:end])
			if status == nil {
				proteins = append(proteins, acid)
				index++
				start += 3
				end += 3
			} else {
				return proteins, nil
			}
		}
		return proteins, nil
	}

}

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
