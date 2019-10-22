package piscine

func Index(s string, toFind string) int {
	sLen := StrLen(s)
	fLen := StrLen(toFind)
	ss := ""
	res := -1
	for i := 0; i < sLen-fLen; i++ {
		for j := 0; j < fLen; j++ {
			ss = ss + string(s[i+j])
		}
		if ss == toFind {
			return i
		}
		ss = ""
	}
	return res
}
