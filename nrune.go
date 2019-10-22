package piscine

func NRune(s string, n int) rune {
	ss := []rune(s)
	return ss[n-1]
}
