package piscine

func IsPrintable(str string) bool {
	for _, a := range str {
		if !(a >= ' ' && a <= '~') {
			return false
		}
	}
	return true
}
