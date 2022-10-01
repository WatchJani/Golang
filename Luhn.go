// Instructions
// Given a number determine whether or not it is valid per the Luhn formula.

// The Luhn algorithm is a simple checksum formula used to validate a variety of identification numbers, such as credit card numbers and Canadian Social Insurance Numbers.

// The task is to check if a given string is valid.

// Validating a Number
// Strings of length 1 or less are not valid. Spaces are allowed in the input, but they should be stripped before checking. All other non-digit characters are disallowed.

// Example 1: valid credit card number
// 4539 3195 0343 6467
// The first step of the Luhn algorithm is to double every second digit, starting from the right. We will be doubling

// 4_3_ 3_9_ 0_4_ 6_6_
// If doubling the number results in a number greater than 9 then subtract 9 from the product. The results of our doubling:

// 8569 6195 0383 3437
// Then sum all of the digits:

// 8+5+6+9+6+1+9+5+0+3+8+3+3+4+3+7 = 80
// If the sum is evenly divisible by 10, then the number is valid. This number is valid!

// Example 2: invalid credit card number
// 8273 1232 7352 0569
// Double the second digits, starting from the right

// 7253 2262 5312 0539
// Sum the digits

// 7+2+5+3+2+2+6+2+5+3+1+2+0+5+3+9 = 57
// 57 is not evenly divisible by 10, so this number is not valid.

package main

import (
	"fmt"
	"strings"
)

// PRVOBITNO JE OVA FUNKCIJA VRAĆALA SLICE TYPE INT, SAD VIŠE NEMA BAŠ MNOGO SMISLA :D
func ParseInt(code string) ([]uint8, bool) {
	var result []uint8
	var my_error bool = true

	for i := 0; i < len(code); i++ {
		if uint8(code[i]) >= '0' && uint8(code[i]) <= '9' {
			result = append(result, uint8(code[i])-'0')
		} else {
			my_error = false
			break
		}
	}

	return result, my_error
}

func Valid(id string) bool {
	id = strings.TrimSpace(strings.ReplaceAll(id, " ", ""))

	my_number, my_error := ParseInt(id)

	if len(id) <= 1 || !my_error {
		return false
	}

	var sum uint8

	for i := len(my_number) - 2; i >= 0; i -= 2 {
		my_number[i] *= 2
		if my_number[i] > 9 {
			my_number[i] -= 9
		}
	}

	for i := 0; i < len(my_number); i++ {
		sum += my_number[i]
	}

	return sum%10 == 0
}

func main() {
	fmt.Println(Valid("    4539 3195 0343 6467"))

	NewString := "test"

	fmt.Printf("%T\n", NewString[0])
	fmt.Printf("%T\n", int(NewString[0]))

}

// ! NAPOMENA KOD OVAKVIH ZADATAKA, STRING NE MOGU PRETVORITI U INT

//NEKO JE URADIO OVAKO TU FUNKCIJU
// s = strings.ReplaceAll(s, " ", "")
// 	if len(s) <= 1 {
// 		return false
// 	}
// 	var sum int
// 	var twice bool
// 	for i := len(s) - 1; i >= 0; i-- {
// 		if !unicode.IsNumber(rune(s[i])) {
// 			return false
// 		}
// 		n := int(s[i] - '0')
// 		if twice { // KAKO JE DOBOR NAPISANO SVAKI DRUGI DA SE GLEDA
// 			n *= 2
// 			if n > 9 {
// 				n -= 9
// 			}
// 			twice = false
// 		} else {
// 			twice = true
// 		}
// 		sum += n
// 	}
// 	return sum%10 == 0
