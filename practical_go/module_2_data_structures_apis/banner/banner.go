package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	banner("Go", 6)
	banner("G♡", 6)

	s := "G♡"
	fmt.Println("len:", len("G")) // 1 byte
	fmt.Println("len:", len(s))   // 7 bytes...

	fmt.Println("s[1]:", s[1])     // 226
	fmt.Printf("s[1]: %c\n", s[1]) // â -> printa apenas um caractere

	// range nas runes...
	for i, c := range s {
		fmt.Printf("%c at %d\n", c, i)
	} // G at 0
	// ❤ at 1
	// at 4
}

// banner("Go", 6)
// Go
// ------

/*
strings sao encodadas em UTF-8
len, s[]: bytes (uin8)
for : runes (int32)
*/

func banner(text string, width int) {
	// BUG: len is in bytes
	// padding := (width - len(text)) / 2
	padding := (width - utf8.RuneCountInString(text)) / 2
	fmt.Print(strings.Repeat(" ", padding))
	fmt.Println(text)
	fmt.Println(strings.Repeat("-", width))
}
