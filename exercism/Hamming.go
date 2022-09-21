package main

import (
	"errors"
	"fmt"
)

func Distance(a, b string) (int, error) {
	number := 0

	if len(a) != len(b) {
		return number, errors.New("Different length")
	}

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			number++
		}
	}

	return number, nil

}

func main() {
	fmt.Println(Distance("GGACTGAAATCTG", "GGACTGAAATCTG"))
}
