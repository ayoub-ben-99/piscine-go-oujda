package piscine

import "github.com/01-edu/z01"

func PrintAll(str string, n bool) {
	for _, r := range str {
		z01.PrintRune(r)
	}
	if n {
		z01.PrintRune('\n')
	}
}
