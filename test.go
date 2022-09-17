package main

import "fmt"

func main() {

	broj := 10

	if broj >= 0 {
		fmt.Println("Vaš broj ", broj, " je pozitivan broj")
	}

	// fmt.Println(broj2 < 0 && broj3 > 2)
	broj /= 5
	broj++
	fmt.Println(broj)

	broj1 := -5
	broj2 := 6.546

	const Pi float64 = 3.14

	fmt.Println(float64(broj1) + broj2*Pi)

	// var broj123 uint8 = 5

	fmt.Printf("%T\n", 0)
	fmt.Printf("%q\n", `asdasd`)

	fmt.Printf("%T\n", 0.23)
}
