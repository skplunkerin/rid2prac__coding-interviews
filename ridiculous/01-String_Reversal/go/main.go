package main

import (
	"fmt"
)

// NOTE: See ../README.md for the problem.

func main() {
	for input, output := range map[string]string{
		"apple":      "elppa",
		"hello":      "olleh",
		"Greetings!": "!sgniteerG",
	} {
		// if reversed := reverse1(input); reversed != output {
		if reversed := reverse2(input); reversed != output {
			panic(fmt.Sprintf(`"%s" isn't the reverse of "%s"; expected "%s"`, reversed, input, output))
		}
	}
	fmt.Println("passed!")
}
