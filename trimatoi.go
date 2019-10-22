package piscine

func TrimAtoi(s string) int {
	n := StrLen(s)
	var list []rune
	symCount := 0
	sign := 1
	isSign := false
	res := 0
	for i := 0; i < n; i++ {
		if (s[i] >= '0' && s[i] <= '9') || ((s[i] == '-' || s[i] == '+') && !isSign && symCount == 0) {
			list = append(list, rune(s[i]))
			symCount++
		}
	}
	tmp := 1
	for i := symCount - 1; i >= 0; i-- {
		if list[i] == '-' {
			sign = -1
			break
		}
		if list[i] == '+' {
			break
		}
		res = res + int(list[i]-48)*tmp
		tmp = tmp * 10
	}
	return res * sign
}
