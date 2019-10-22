package piscine

func BasicJoin(strs []string) string {
	res := ""
	for _, a := range strs {
		res = res + a
	}
	return res
}
