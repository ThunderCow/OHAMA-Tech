package main

import (
	"fmt"
)

func main() {

	fmt.Println("Calling the function to show the three different forms")
	iterateOverExtendedASCIIStringLiteral()
	cp := "\x22\x20\x80\x20\xf7\x20\xbe\x20\x64\x6f\x6c\x6c\x61\x72\x20\x22"
	str := ExtendedASCIIText(cp)
	fmt.Printf("%s", str,)


}

//Function for iteration of ASCII Codes and print out in three different forms.
func iterateOverExtendedASCIIStringLiteral(){

	for i := 128 ; i <= 255 ; i++ {

		// This prints hexadecimal value
		fmt.Printf("%X\t", i,)

		// This prints the symbol
		fmt.Printf("%c\t", i,)

		// This prints the binary value
		fmt.Printf("%b\n", i,)
	}
}


// Code to solve the Oppgave 4b

func ExtendedASCIIText(cp string) string {
	r := make([]rune, len(cp))
	for i := 0; i < len(cp); i++ {
		r[i] = ascii[cp[i]]
	}
	return string(r)
}

// Function to check the values and set the condition.
func init() {
	for i, r := range ascii {
		if r == 0 {
			ascii[i] = rune(i)
		}
	}
}

// cp1252 to Unicode table, this table shows the representation of unicode alongwith other formats.
// ftp://ftp.unicode.org/Public/MAPPINGS/VENDORS/MICSFT/WINDOWS/CP1252.TXT
var ascii = [256]rune{
	0x80: '\u20AC', // EURO SIGN
	0x81: '\uFFFD', // UNDEFINED
	0x82: '\u201A', // SINGLE LOW-9 QUOTATION MARK
	0x83: '\u0192', // LATIN SMALL LETTER F WITH HOOK
	0x84: '\u201E', // DOUBLE LOW-9 QUOTATION MARK
	0x85: '\u2026', // HORIZONTAL ELLIPSIS
	0x86: '\u2020', // DAGGER
	0x87: '\u2021', // DOUBLE DAGGER
	0x88: '\u02C6', // MODIFIER LETTER CIRCUMFLEX ACCENT
	0x89: '\u2030', // PER MILLE SIGN
	0x8A: '\u0160', // LATIN CAPITAL LETTER S WITH CARON
	0x8B: '\u2039', // SINGLE LEFT-POINTING ANGLE QUOTATION MARK
	0x8C: '\u0152', // LATIN CAPITAL LIGATURE OE
	0x8D: '\uFFFD', // UNDEFINED
	0x8E: '\u017D', // LATIN CAPITAL LETTER Z WITH CARON
	0x8F: '\uFFFD', // UNDEFINED
	0x90: '\uFFFD', // UNDEFINED
	0x91: '\u2018', // LEFT SINGLE QUOTATION MARK
	0x92: '\u2019', // RIGHT SINGLE QUOTATION MARK
	0x93: '\u201C', // LEFT DOUBLE QUOTATION MARK
	0x94: '\u201D', // RIGHT DOUBLE QUOTATION MARK
	0x95: '\u2022', // BULLET
	0x96: '\u2013', // EN DASH
	0x97: '\u2014', // EM DASH
	0x98: '\u02DC', // SMALL TILDE
	0x99: '\u2122', // TRADE MARK SIGN
	0x9A: '\u0161', // LATIN SMALL LETTER S WITH CARON
	0x9B: '\u203A', // SINGLE RIGHT-POINTING ANGLE QUOTATION MARK
	0x9C: '\u0153', // LATIN SMALL LIGATURE OE
	0x9D: '\uFFFD', // UNDEFINED
	0x9E: '\u017E', // LATIN SMALL LETTER Z WITH CARON
	0x9F: '\u0178', // LATIN CAPITAL LETTER Y WITH DIAERESIS
}
