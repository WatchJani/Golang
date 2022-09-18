package main

import "fmt"

type Trougao struct {
	a int
	b int
	c int
}

func Test(trougao Trougao) {
	fmt.Println(trougao)
}

func main() {
	trougao := Trougao{5, 4, 8}

	// fmt.Println(trougao)

	// number := 2

	Test(trougao)
}
