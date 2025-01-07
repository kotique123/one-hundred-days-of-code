// Package proverb is a small library that implements proverb generation.
package proverb

/*
goos: windows
goarch: amd64
pkg: proverb
cpu: 13th Gen Intel(R) Core(TM) i7-13700KF
BenchmarkProverb
BenchmarkProverb-24      1982730               599.1 ns/op          1008 B/op         22 allocs/op
*/
// Proverb returns a slice of strings representing a proverb generated based on
// a set of inputs.
func Proverb(rhyme []string) []string {
	// Initializing function-wide(local) variables
	pre := "For want of a "              // 1st part of a proverb
	last := "And all for the want of a " // last part of a proverb
	// if length of an argument is one, we want to return just the last part
	// of a proverb.
	if len(rhyme) == 1 {
		proverb := []string{(last + rhyme[0] + ".")}
		return proverb
	} else if len(rhyme) > 1 {
		// initalize a non empty slince in which we will write data based on
		// its index. I found out that this solution is more performant than
		// using append keyword or using print funcitons from fmt package.
		proverb := make([]string, len(rhyme))
		for i := 0; i < len(rhyme)-1; i++ {
			proverb[i] = (pre + rhyme[i] + " the " + rhyme[i+1] + " was lost.")
		}
		// then we add the last element of a proverb.
		proverb[len(rhyme)-1] = (last + rhyme[0] + ".")
		return proverb
	}
	// in case of length being 0, we return an empty slice.
	return []string{}
}
