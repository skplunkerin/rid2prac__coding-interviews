package main

import (
	"io/ioutil"
	"os"
)

func main() {
	for _, fileName := range []string{
		"01-String Reversal",
		"02-Palindromes",
		"03-Integer Reversal",
		"04-MaxChars",
		"05-FizzBuzz",
		"06-Array Chunking",
		"07-Anagrams",
		"08-Sentence Capitalization",
		"09-Printing Steps",
		"10-Find The Vowels",
		"11-Matrix Spiral",
		"12-Runtime Complexity",
		"13-Fibonacci",
		"14-Queues",
		"15-Im Weaving You Behind",
		"16-Stack Data Structures",
		"17-Two Becoming One",
		"18-Linked Lists",
		"19-Finding the Midpoint",
		"20-Circular Lists",
		"21-Stepping back from the tail",
		"22-Building a Tree",
		"23-Tree width and level width",
		"24-Binary Search Trees",
		"25-Validating Binary Search Trees",
		"26-Events",
		"27-Design - Build a Twitter",
		"28-Bubble Sorting",
		"29-Selection Sorting",
		"30-Merge Sort",
	} {
		// Create the directory in 2 locations
		for _, dirPath := range []string{
			"ridiculous/",
			"practical/",
		} {
			if err := os.Mkdir(dirPath+fileName, 0755); err != nil {
				panic(err)
			}
			// Create the blank .gitkeep file
			if err := createEmptyFile(dirPath + fileName + "/.gitkeep"); err != nil {
				panic(err)
			}
		}
	}
}

func createEmptyFile(name string) error {
	d := []byte("")
	return ioutil.WriteFile(name, d, 0644)
}
