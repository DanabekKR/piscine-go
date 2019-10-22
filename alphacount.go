package piscine

func AlphaCount(str string) int {
	res := 0
	for i := 0; i < StrLen(str); i++ {
		if (str[i] >= 'a' && str[i] <= 'z') || (str[i] >= 'A' && str[i] <= 'Z') {
			res++
		}
	}
	return res
}
