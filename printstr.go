package piscine

import (
	"github.com/01-edu/z01"
)

func PrintStr(str string) {
	newStr := []rune(str)
	for _, sym := range newStr {
		z01.PrintRune(sym)
	}
}
