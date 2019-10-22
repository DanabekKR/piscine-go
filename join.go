package piscine

func Join(strs []string, sep string) string {
	res := ""
	n := 0
	for range strs {
		n++
	}
	for i, a := range strs {
		if i == (n - 1) {
			res = res + a
		} else {
			res = res + a + sep
		}
	}
	return res
}
