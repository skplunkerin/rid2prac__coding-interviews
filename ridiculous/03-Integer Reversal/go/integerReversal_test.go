package main

import (
	"fmt"
	"testing"
)

func TestIntegerReversal(t *testing.T) {
	testInts := map[int]int{
		15:  51,
		981: 189,
		500: 5,
		-15: -51,
		-90: -9,
	}
	tests := map[string]struct {
		fn func(int) int
	}{
		"success": {
			fn: func(i int) int {
				return integerReverse(i)
			},
		},
	}
	for name, test := range tests {
		for input, output := range testInts {
			t.Run(fmt.Sprintf("%s_%d", name, input), func(t *testing.T) {
				if out := test.fn(input); out != output {
					t.Fatalf("got %d, expected %d", out, output)
				}
			})
		}
	}
}
