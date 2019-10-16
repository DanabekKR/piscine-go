package piscine

func StrLen(str string) int {
	newStr := []rune(str)
	count := 0
	for range newStr {
		count++
	}
	return count
}
