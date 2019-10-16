package piscine

func StrRev(str string) string {
	newStr := []rune(str)
	count := StrLen(str)
	for i := 0; i < (count / 2); i++ {
		sym := newStr[i]
		newStr[i] = newStr[count-i-1]
		newStr[count-i-1] = sym
	}
	return string(newStr)
}
