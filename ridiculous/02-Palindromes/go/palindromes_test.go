package main

import (
	"testing"
)

func TestPalindromes(t *testing.T) {
	testStrings := map[string]bool{
		"abba":    true,
		"racecar": true,
		"cat tac": true,
		"abcdefg": false,
		"bròwn":   false,
	}
	// bonusStrings tuna is bacon
	bonusStrings := map[string]bool{
		"as⃝df̅ds⃝a": true,
		"as⃝d ds⃝a":  true,
		"as⃝df̅":     false,
	}
	tests := map[string]struct {
		fn                    func(string) bool
		supportsCombinedChars bool
	}{
		"simple": {
			supportsCombinedChars: true,
			fn: func(s string) bool {
				return palindromeSimple(s)
			},
		},
	}
	for name, test := range tests {
		for input, expectedResult := range testStrings {
			t.Run(name+"-"+input, func(t *testing.T) {
				isPalindrome := test.fn(input)
				if isPalindrome != expectedResult {
					t.Fatalf("got %t, expected %t", isPalindrome, expectedResult)
				}
			})
		}
		if test.supportsCombinedChars {
			for input, expectedResult := range bonusStrings {
				t.Run(name+"-"+input, func(t *testing.T) {
					isPalindrome := test.fn(input)
					if isPalindrome != expectedResult {
						t.Fatalf("got %t, expected %t", isPalindrome, expectedResult)
					}
				})
			}
		}
	}
}
