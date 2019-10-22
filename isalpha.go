package piscine

func IsAlpha(str string) bool {
	for _, a := range str {
		if !isLetter(a) {
			return false
		}
	}
	return true
}
