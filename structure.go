package main

import "fmt"

type Trougao struct {
	a int
	b int
	c int
}

type Price int

func main() {
	trougao1 := &Trougao{1, 2, 3}
	trougao2 := &Trougao{}

	trougao2.a = 15

	fmt.Println(trougao2.a)
	fmt.Println(trougao1.a)

	var trougao3 Trougao
	trougao3 = Trougao{54, 654, 4}

	var p *Trougao = &trougao3 //????

	fmt.Println(p)

	fmt.Println(&trougao3)

	cijena := Price(5000)

	cijena2 := uint8(1)
	fmt.Printf("\n\n vrijednost: %v Type: %T \n\n", cijena, cijena2)

	fmt.Printf("%T\n", 65454.654546)
	fmt.Printf("%T\n", 1)
	fmt.Println(cijena)
}
