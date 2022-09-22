package main

import "fmt"

func main() {

	array := [...]int{1, 2, 3}
	slice := []int{123, 132, 321}
	maps := map[int]int{
		1: 5,
		2: 3,
		4: 2,
	}

	fmt.Println(array[0])
	fmt.Println(slice[0])
	fmt.Println(maps[1])

	slice = append(slice, 321)
	slice[0] = 0

	fmt.Println(slice)
	fmt.Printf("%v %v %v \n\n", len(slice), cap(slice), slice)

	slice = append(slice[:2], slice[3:]...)
	fmt.Println(slice)
	
}
