package piscine

func BasicAtoi(s string) int {
	count := StrLen(s)
	newStr := []rune(s)
	val := 0
	for i := 0; i < count; i++ {
		t := 1
		for j := 0; j < count-i-1; j++ {
			t = t * 10
		}
		sym := 0
		for k := '0'; k < newStr[i]; k++ {
			sym++
		}

		val = val + sym*t
	}
	return val
}
