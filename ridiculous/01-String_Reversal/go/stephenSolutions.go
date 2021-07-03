package main

import "strings"

// NOTE: The solutions in here are from Stephen Grider's course, check it out
// to support Stephen and have him teach you how to improve/be ready for your
// next software engineer interview:
// [The Coding Interview Bootcamp: Algorithms + Data Structures](https://www.udemy.com/share/101WNk2@PUdgVFpSSFALckNGO0hOVD5u/)

// reverse2 takes the string and returns the string in reversed order.
//
// This was Stephen's most elegant approach (IMO).
func reverse2(s string) (reversed string) {
	// convert the input string to []string:
	slice := strings.Split(s, "")
	// create a new string by looping through slice and adding each character to
	// the beginning of reversed:
	for _, char := range slice {
		reversed = char + reversed
	}
	return reversed
}

// reverse3 takes the string and returns the string in reversed order.
func reverse3(s string) (reversed string) {
	return reversed
}

// reverse4 takes the string and returns the string in reversed order.
func reverse4(s string) (reversed string) {
	return reversed
}
