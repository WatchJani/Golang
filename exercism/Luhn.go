package main

import (
	"fmt"
	"strings"
)

func ParseInt(code string) bool {

	return true
}

func Valid(id string) bool {
	id = strings.TrimSpace(strings.ReplaceAll(id, " ", ""))

	// opara := rune(id)

	// if len(id) <= 1 {
	// 	return false
	// }

	// if secure{
	// 	return false
	// }

	// rune := []uint{'1', '4'}

	// fmt.Printf("%T \n", opara)

	return true
}

func main() {
	fmt.Println(Valid("234 567 891 234"))

	my_string := "23456791234"

	runes := []rune(my_string)

	var result []int

	for i := 0; i < len(runes); i++ {
		result = append(result, int(runes[i]))
	}

	fmt.Printf("%T", result)
}
