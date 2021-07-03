package main

import "strings"

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
