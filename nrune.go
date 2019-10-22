package piscine

func NRune(s string, n int) rune {
	for i, a := range s {
		if i+1 == n {
			return a
		}
	}
	return 0
}
