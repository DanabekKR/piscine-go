package main

import (
	"os"

	"github.com/01-edu/z01"
)

func printStr(str string) {
	arrayStr := []rune(str)

	for i := 0; i < len(arrayStr); i++ {
		z01.PrintRune(arrayStr[i])
	}
	z01.PrintRune('\n')
}

func isEven(nbr int) bool {
	if nbr%2 == 1 {
		return false
	}
	return true

}

func main() {
	args := os.Args
	lengthOfArg := 0
	for range args {
		lengthOfArg++
	}
	evenMsg := "I have an even number of arguments"
	oddMsg := "I have an odd number of arguments"
	if isEven(lengthOfArg + 1) {
		printStr(evenMsg)
	} else {
		printStr(oddMsg)
	}
}
