package piscine

func Join(strs []string, sep string) string {
	res := ""
	for _, a := range strs {
		res = res + a + sep
	}
	return res
}
