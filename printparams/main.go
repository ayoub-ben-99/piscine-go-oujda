package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	for _, v := range os.Args[1:] {
		for _, t := range v {
			z01.PrintRune(t)
		}
		z01.PrintRune('\n')
	}
}
