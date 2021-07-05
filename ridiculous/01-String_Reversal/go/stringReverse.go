package main

import (
	"strings"
	"unicode"
	"unicode/utf8"
)

// NOTE: See ../README.md for the problem.

// =================================================
// NOTE: The solutions in here are my own solutions.
// reverse1 takes the string and returns the string in reversed order.
//
// NOTE: This was my first attempt; it works but isn't as simple as it could be.
func reverse1(s string) (reversed string) {
	// convert the input string to []string:
	slice := strings.Split(s, "")
	// create a new string by looping through slice in reverse order:
	for i := range slice {
		reversed += slice[len(slice)-i-1]
	}
	return reversed
}

// reverse2 takes the string and returns the string in reversed order.
//
// I forgot that strings can be treated like an slice.
func reverse2(s string) (reversed string) {
	// create a new string by looping through each rune in the string:
	for _, r := range s {
		reversed = string(r) + reversed
	}
	return reversed
}

// reverse3 takes the string and returns the string in reversed order.
//
// This handles runes in a string that are more than 1 bit in size.
// NOTE: this is taken from stringutil.Reverse(), how it works is it swaps
// places with each []rune position; first and last swap, second and second
// last, etc.
func reverse3(s string) string {
	// return s
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

// reverse4 interprets its argument as UTF-8 and ignores bytes that do not form
// valid UTF-8. Return value is UTF-8.
//
// sources:
// - https://stackoverflow.com/a/30973997/1180523
// - https://rosettacode.org/wiki/Reverse_a_string#Go
func reverse4(s string) string {
	if s == "" {
		return ""
	}
	p := []rune(s)
	r := make([]rune, len(p))
	start := len(r)
	for i := 0; i < len(p); {
		// quietly skip invalid UTF-8
		if p[i] == utf8.RuneError {
			i++
			continue
		}
		j := i + 1
		for j < len(p) && (unicode.Is(unicode.Mn, p[j]) ||
			unicode.Is(unicode.Me, p[j]) || unicode.Is(unicode.Mc, p[j])) {
			j++
		}
		for k := j - 1; k >= i; k-- {
			start--
			r[start] = p[k]
		}
		i = j
	}
	return (string(r[start:]))
}

// /end of my own solutions
// =================================================

// =================================================
// NOTE: The solutions in here are from Stephen Grider's course (javascript
// based), check it out to support Stephen and have him teach you how to
// improve/be ready for your next software engineer interview:
// [The Coding Interview Bootcamp: Algorithms + Data Structures](https://www.udemy.com/share/101WNk2@PUdgVFpSSFALckNGO0hOVD5u/)

// reverse5 takes the string and returns the string in reversed order.
//
// This was Stephen's most elegant approach (IMO).
func reverse5(s string) (reversed string) {
	// convert the input string to []string:
	slice := strings.Split(s, "")
	// create a new string by looping through slice and adding each character to
	// the beginning of reversed:
	for _, char := range slice {
		reversed = char + reversed
	}
	return reversed
}

// /end of Stephen Grider's course solutions (javascript based)
// =================================================

// =================================================
// theTruestReverse handles both UTF8 and Combined Characters, correctly
// reversing any string.
//
// https://stackoverflow.com/a/44350406/1180523
func theTruestReverse(text string) string {
	// reverse is a simple replacement of positions from first/last,
	// second/second-last, etc in a []rune.
	reverse := func(runes []rune) {
		length := len(runes)
		for i, j := 0, length-1; i < length/2; i, j = i+1, j-1 {
			runes[i], runes[j] = runes[j], runes[i]
		}
	}
	// isMark determines whether the rune is a marker.
	isMark := func(r rune) bool {
		return unicode.Is(unicode.Mn, r) ||
			unicode.Is(unicode.Me, r) ||
			unicode.Is(unicode.Mc, r)
	}

	// convert string to []rune
	textRunes := []rune(text)
	textRunesLength := len(textRunes)

	// return text if it's ""
	if textRunesLength <= 1 {
		return text
	}

	// NOTE: what characters are a marker?
	// for _, jj := range textRunes {
	// 	if isMark(jj) {
	// 		fmt.Printf(`"%s" is a marker`+"\n", string(jj))
	// 	}
	// }

	// loop through []rune to search for combined characters
	i, j := 0, 0
	for i < textRunesLength && j < textRunesLength {
		// keep incrementing j until a marker is found
		j = i + 1
		for j < textRunesLength && isMark(textRunes[j]) {
			j++
		}
		// if a marker was found
		if isMark(textRunes[j-1]) {
			// reverse the combined characters
			reverse(textRunes[i:j])
		}

		i = j
	}

	// reverse the entire array, this now works since each combined
	// character has already been reversed.
	reverse(textRunes)

	return string(textRunes)
}

// /combinedChars support
// =================================================
