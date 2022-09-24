package main

import (
	"errors"
	"fmt"
	"math"
)

func Square(number int) (uint64, error) {

	if number <= 0 || number > 65 {
		return 0, errors.New("error")
	}

	return uint64(math.Pow(2, float64(number-1))), nil
}

func Total() uint64 {
	return 18446744073709551615  //Traze najbrzi algoritam 
}

func main() {
	fmt.Println(Total())
}
