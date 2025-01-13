package acronym

import (
	"strings"
)

func Abbreviate(s string) string {
	s = strings.ReplaceAll(s, "-", " ")
	s = strings.ReplaceAll(s, "_", " ")
	s = strings.ToUpper(s)

	words := strings.Fields(s)

	acronym := make([]byte, 0, len(words))
	for _, word := range words {
		acronym = append(acronym, word[0])
	}
	return string(acronym)
}
