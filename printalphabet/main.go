package main

import "github.com/01-edu/z01"
func main(){
	for i:=0; i<26; i++{
		z01.PrintRune(97+i)
	}
	z01.PrintRune('\n')
}