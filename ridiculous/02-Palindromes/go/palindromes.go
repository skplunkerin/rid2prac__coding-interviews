package main

import (
	"unicode"
)

// palindromeSimple reverses a string (taking into account multiple
// codepoints/runes) and compares it with the original string, returning true if
// it matches.
//
// NOTE: Palindromes are strings that are the same string when reversed.
func palindromeSimple(s string) bool {
	// before reversing and comparing, loop through the string to reverse any
	// multiple codepoints first, then compare the reversed of that string.
	r := []rune(s)

	// loop through the []rune to find and replace any of the special characters.
	// https://stackoverflow.com/a/44350406/1180523
	for i, j := 0, 0; i < len(r) && j < len(r); {
		// begin loop with j being ahead of i
		j = i + 1
		// if a marked character is found, reverse
		if j < len(r) && isMarkedChar(r[j]) {
			j++
			reverse(r[i:j])
		}
		// end loop with i & j having the same value
		i = j
	}

	reverse(r)
	return string(r) == s
}

// isMarkedChar returns true if the passed in rune is a "marked" character.
//
// NOTE: Only checking against unicode.Mark covers the individual checking of Mc,
// Me, and Mn.
func isMarkedChar(r rune) bool {
	return unicode.Is(unicode.Mark, r)
}

// reverse reverses the order of []rune.
func reverse(r []rune) {
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
}
