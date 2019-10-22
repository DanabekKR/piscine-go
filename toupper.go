package piscine

func ToUpper(s string) string {

	res := ""
	for _, a := range s {
		if a >= 'a' && a <= 'z' {
			a = a - 32
		}
		res = res + string(rune(a))
	}
	return res
}
