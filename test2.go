package main

import "fmt"

func main() {
	if 1 == 1 {
		fmt.Println("radi")
	}

	const RealPi float64 = 3.14

	fmt.Printf("%T\n", 3.14)

	if Pi := 3.14; RealPi == Pi {
		fmt.Println("Stvarno se poduraraju")
	}

	if novac, cijena, plation := 1000, 5000, true; plation && novac-cijena > 0 {
		fmt.Println(novac)
	} else {
		fmt.Println("Nemate dovoljno na računu da bi ovo platili :D")
	}
}
