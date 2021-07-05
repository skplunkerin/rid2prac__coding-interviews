package main

import (
	"testing"
)

func TestReverse(t *testing.T) {
	// testStrings contains a slice of {"inputString": "expectedOutputString"}.
	testStrings := map[string]string{
		"apple":      "elppa",
		"hello":      "olleh",
		"Greetings!": "!sgniteerG",
		"bròwn":      "nwòrb",
	}
	// combinedChars are Unicode strings with combining characters, the majority of
	// reverse string solutions (at least in Go) don't support these types of
	// characters correctly.
	// https://stackoverflow.com/a/30973997/1180523
	// https://rosettacode.org/wiki/Reverse_a_string#Go
	combinedChars := map[string]string{
		"👯‍♂️🙅‍♂️": "♂️‍🙅♂️‍👯",
		"as⃝df̅":   "f̅ds⃝a",
	}

	tests := map[string]struct {
		function              func(string) string
		supportsCombinedChars bool
	}{
		"reverse1": {
			supportsCombinedChars: false,
			function: func(i string) string {
				return reverse1(i)
			},
		},
		"reverse2": {
			supportsCombinedChars: false,
			function: func(i string) string {
				return reverse2(i)
			},
		},
		"reverse3": {
			supportsCombinedChars: false,
			function: func(i string) string {
				return reverse3(i)
			},
		},
		"reverse4": {
			supportsCombinedChars: false,
			function: func(i string) string {
				return reverse4(i)
			},
		},
		"reverse5": {
			supportsCombinedChars: false,
			function: func(i string) string {
				return reverse5(i)
			},
		},
		"the only true solution": {
			supportsCombinedChars: true,
			function: func(i string) string {
				return theTruestReverse(i)
			},
		},
	}
	for name, test := range tests {
		for input, expectedReversed := range testStrings {
			t.Run(name+"_"+input, func(t *testing.T) {
				reversed := test.function(input)
				if reversed != expectedReversed {
					t.Fatalf(`got "%s", expected "%s"`, reversed, expectedReversed)
				}
			})
		}
		if test.supportsCombinedChars {
			for input, expectedReversed := range combinedChars {
				t.Run(name+"-combinedChars-"+input, func(t *testing.T) {
					reversed := test.function(input)
					if reversed != expectedReversed {
						t.Fatalf(`got "%s", expected "%s"`, reversed, expectedReversed)
					}
				})
			}
		}
	}
}
