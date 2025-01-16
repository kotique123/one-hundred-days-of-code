package series

func All(n int, s string) []string {
	slice := make([]string, 0, len(s)-n+1)
	for i := 0; i < len(s)-n+1; i++ {
		slice = append(slice, s[i:i+n])
	}
	return slice
}

func UnsafeFirst(n int, s string) string {
	return s[:n]
}

func First(n int, s string) (first string, ok bool) {
	if n > len(s) {
		return "", false
	}
	return s[:n], true
}
