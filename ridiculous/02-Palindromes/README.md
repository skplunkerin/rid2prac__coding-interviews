# Palindromes

**categories:** [tricky](https://github.com/skplunkerin/rid2prac__coding-interviews#tricky)

---

## Problem:

Given a string, return true if it's a palindrome or false if it is not.
Palindromes are strings that form the same word if it is reversed.

_**DO** include spaces, numbers, and punctuations in strings._

### Examples:

| input       | expected output |
| ----------- | --------------- |
| `"abba"`    | `true`          |
| `"racecar"` | `true`          |
| `"cat tac"` | `true`          |
| `"abcdefg"` | `false`         |
| `"bròwn"`   | `false`         |

## Extra Credit:

Given a string with [multiple codepoints/runes](https://www.reedbeta.com/blog/programmers-intro-to-unicode/#combining-marks), return true if it's a palindrome
or false if it is not. Palindromes are strings that form the same word if it is
reversed.

### Examples:

| input         | expected output |
| ------------- | --------------- |
| `"as⃝df̅ds⃝a"` | `true`          |
| `"as⃝d ds⃝a"` | `true`          |
| `"as⃝df̅"`     | `false`         |
