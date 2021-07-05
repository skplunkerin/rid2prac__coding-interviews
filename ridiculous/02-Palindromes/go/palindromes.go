package main

import "unicode"

// palindromeSimple reverses a string (not taking into account multiple
// codepoints/runes) and compares it with the original string, returning true if
// it matches.
//
// NOTE: Palindromes are strings that are the same when reversed.
func palindromeSimple(s string) bool {
	// before reversing and comparing, loop through the string to reverse any
	// multiple codepoints first, then compare the reversed of that string.
	r := []rune(s)

	// loop through the []rune to find and replace any of the special characters.
	for i := 0; i < len(r); i++ {
		// TODO: fix this, it's not working right
		if isMarkedChar(r[i]) {
			r[i], r[i+1] = r[i+1], r[i]
			i = i + 2
		}
	}

	return reverse(string(r)) == s
}

// isMarkedChar returns true if the passed in rune is a "marked" character.
func isMarkedChar(r rune) bool {
	// TODO: is checking againt unicode.Mark enough? Or will I need to check
	// against all Mark values (Mc, Me, and Mn)?
	return unicode.Is(unicode.Mark, r)
}

func reverse(s string) string {
	r := []rune(s)
	for i, l := 0, len(r)-1; i < len(r)/2; i, l = i+1, l-1 {
		r[i], r[l] = r[l], r[i]
	}
	return string(r)
}
