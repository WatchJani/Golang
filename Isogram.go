// Instructions
// Determine if a word or phrase is an isogram.

// An isogram (also known as a "non-pattern word") is a word or phrase without a repeating letter, however spaces and hyphens are allowed to appear multiple times.

// Examples of isograms:

// lumberjacks
// background
// downstream
// six-year-old
// The word isograms, however, is not an isogram, because the s repeats.

package main

import (
	"fmt"
	"unicode"
)

func IsIsogram(word string) bool {
	char := make(map[rune]int)

	for _, rune := range word {
		_, exist := char[unicode.ToUpper(rune)]
		if exist {
			return false
		}
		if unicode.IsLetter(rune) {
			char[unicode.ToUpper(rune)] += 1
		}
	}

	fmt.Println(char)

	return true
}

func main() {
	fmt.Println(IsIsogram("janko---b  SD"))
}
