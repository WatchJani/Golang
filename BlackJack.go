package main

import "fmt"

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	switch card {
	case "ace":
		return 11
	case "two":
		return 2
	case "three":
		return 3
	case "four":
		return 4
	case "five":
		return 5
	case "six":
		return 6
	case "seven":
		return 7
	case "eight":
		return 8
	case "nine":
		return 9
	case "ten", "jack", "queen", "king":
		return 10
	default:
		return 0
	}

	//PUNO BOLJE RJESENJE JE KORIŠTENJE MAPE
}
func FirstTurn(card1, card2, dealerCard string) string {
	MyDilerCard := ParseCard(dealerCard)
	switch sum := ParseCard(card1) + ParseCard(card2); {
	case sum == 22:
		return "P"
	case sum == 21:
		if MyDilerCard >= 10 {
			return "S"
		} else {
			return "W"
		}
	case sum >= 17 && sum <= 20:
		return "S"
	case sum >= 12 && sum <= 16:
		if MyDilerCard >= 7 {
			return "H"
		} else {
			return "S"
		}
	default:
		return "H"
	}
}

func main() {
	fmt.Println(FirstTurn("king", "ace", "queen"))
}
