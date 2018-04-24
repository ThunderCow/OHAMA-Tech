package main

import (
	"fmt"
	"unicode"
)

func main() {
	fmt.Printf("len([]rune(s)) = %d, len([]byte(s)) = %d\n", len([]rune(s)), len([]byte(s)))
	fmt.Println(IsAsciiPrintable(s), IsAsciiPrintable("test"))
}

// checks if s is ascii and printable, aka doesn't include tab, backspace, etc.
const s = "main.ExtendedASCIIText()"
func IsAsciiPrintable(s string) bool {
	for _, r := range s {
		if r > unicode.MaxASCII || !unicode.IsPrint(r) {
			return false
		}
	}
	return true
}

