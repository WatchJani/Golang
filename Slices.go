package main

import "fmt"

func main() {

	var intSlice []int
	var stringSlice []string

	fmt.Printf("%v %v\n\n", intSlice, stringSlice)

	intSlice = make([]int, 0)
	stringSlice = make([]string, 10, 20)

	intSlice = []int{2}

	fmt.Printf("intSlice \tLen: %v \tCap: %v\n", len(intSlice), cap(intSlice))
	fmt.Printf("intSlice \tLen: %v \tCap: %v\n", len(stringSlice), cap(stringSlice))

}
