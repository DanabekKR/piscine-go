package piscine

func SplitWhiteSpaces(str string) []string {
	n := 0
	pos := 0
	str = str + " "
	for i, a := range str {
		if a == ' ' {
			tmp := str[pos:i]
			if tmp != "" {
				n++
			}
			pos = i + 1
		}
	}
	res := make([]string, n)
	k := 0
	ppos := 0
	for i, a := range str {
		if a == ' ' {
			tmp := str[ppos:i]
			if tmp != "" {
				res[k] = tmp
				k++
			}
			ppos = i + 1
		}
	}
	return res
}
