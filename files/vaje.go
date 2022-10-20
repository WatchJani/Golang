package main

import (
	"fmt"
	"math"
)

func PrimeNumber(number int) bool {
	counter := 0

	for i := 1; i <= int(math.Sqrt(float64(number))); i++ {
		if number%i == 0 {
			counter++
			if counter > 1 {
				return false
			}
		}
	}

	return true
}

func main() {
	for i := 2; i < 10000000; i++ {
		if PrimeNumber(i) {
			fmt.Println(i)
		}
	}

	// fmt.Println(PrimeNumber(1))

}
