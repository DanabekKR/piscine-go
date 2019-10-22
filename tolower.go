package piscine

func ToLower(s string) string {

	res := ""
	for _, a := range s {
		if a >= 'A' && a <= 'Z' {
			a = a + 32
		}
		res = res + string(rune(a))
	}
	return res
}
