// Instructions
// Given a word, compute the Scrabble score for that word.

// Letter Values
// You'll need these:

// Letter                           Value
// A, E, I, O, U, L, N, R, S, T       1
// D, G                               2
// B, C, M, P                         3
// F, H, V, W, Y                      4
// K                                  5
// J, X                               8
// Q, Z                               10
// Examples
// "cabbage" should be scored as worth 14 points:

// 3 points for C
// 1 point for A, twice
// 3 points for B, twice
// 2 points for G
// 1 point for E
// And to total:

// 3 + 2*1 + 2*3 + 2 + 1
// = 3 + 2 + 6 + 3
// = 5 + 9
// = 14
// Extensions
// You can play a double or a triple letter.
// You can play a double or a triple word.

package main

import (
	"fmt"
	"unicode"
)

//strings za velika mala slova

func Score(word string) int {
	var Scrabble_Score int

	for _, rune := range word {
		switch unicode.ToUpper(rune) {
		case 'A', 'E', 'I', 'O', 'U', 'L', 'N', 'R', 'S', 'T':
			Scrabble_Score += 1
		case 'D', 'G':
			Scrabble_Score += 2
		case 'B', 'C', 'M', 'P':
			Scrabble_Score += 3
		case 'F', 'H', 'V', 'W', 'Y':
			Scrabble_Score += 4
		case 'K':
			Scrabble_Score += 5
		case 'J', 'X':
			Scrabble_Score += 8
		case 'Q', 'Z':
			Scrabble_Score += 10
		}
	}

	return Scrabble_Score
}

func main() {
	fmt.Println(Score("zoo"))
}
