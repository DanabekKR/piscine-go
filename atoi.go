package piscine

func Atoi(s string) int {
	newStr := []rune(s)
	count := StrLen(s)
	sstr := ""
	for i := 1; i < count; i++ {
		sstr = sstr + string(newStr[i])
	}
	if newStr[0] == '-' {
		return -1 * BasicAtoi2(sstr)
	}
	if s == "+" || s == "-" {
		return 0
	}
	if newStr[0] == '+' {
		return BasicAtoi2(sstr)
	}
	return BasicAtoi2(s)
}
