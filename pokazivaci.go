package main

import "fmt"

func main() {
	// fmt.Printf("%6d", 12)

	number := 21
	p := &number

	*p = 100 //cool

	fmt.Println(*p, &number, number)
}
