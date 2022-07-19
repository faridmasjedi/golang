package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	const s = "سلام"

	// length of the raw bytes
	fmt.Println("len:", len(s))

	// raw byte values at each index
	// %x : hex values
	for i := 0; i < len(s); i++ {
		fmt.Printf("i : %x \n", s[i])
	}

	fmt.Println()

	// This will return the number of runes in the string
	fmt.Println("Rune count:", utf8.RuneCountInString(s))

	fmt.Println()
	// %#U will return the string indexed
	// %d will return the index
	for i, v := range s {
		fmt.Printf("%#U starts at %d\n", v, i)
	}

	fmt.Println()

	for i, w := 0, 0; i < len(s); i += w {
		runeValue, width := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%#U starts at %d\n", runeValue, i)
		w = width

		examineRune(runeValue)
	}
}

func examineRune(r rune) {
	if r == 't' {
		fmt.Println("found t")
	} else if r == 'ل' {
		fmt.Println("found ل")
	}
}
