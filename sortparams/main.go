package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args
	count := 0
	for range arguments {
		count++
	}
	for i := 1; i < count; i++ {
		for j := i; j < count; j++ {
			if arguments[i] > arguments[j] {
				tmp := arguments[i]
				arguments[i] = arguments[j]
				arguments[j] = tmp
			}
		}

	}
	for i := 1; i < count; i++ {
		pn := arguments[i]
		if i > 0 {
			for _, a := range pn {
				z01.PrintRune(a)
			}
		}
		z01.PrintRune('\n')
	}
}
