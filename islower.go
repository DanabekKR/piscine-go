package piscine

func IsLower(str string) bool {
	for _, a := range str {
		if !(a >= 'a' && a <= 'z') {
			return false
		}
	}
	return true
}
