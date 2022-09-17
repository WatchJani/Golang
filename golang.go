package main

import "fmt"

func main() {
	// zivotinja := "medo"
	// var Ucenik string = "Janko"
	var Age uint8 = 2
	// var Ime string
	// Ime = "Janko"
	ostatak := 0.151651
	fmt.Println(Age + uint8(ostatak))

	ime, broj := "Janko", 18

	fmt.Println(ime, broj)

	type User struct {
		UserName  string
		ID        int
		Available bool
	}

	var coek User

	coek = User{"Janko", 22121321, true}
	coek2 := User{"Pero", 654654, false}

	fmt.Println(coek, coek2)

	const Pi = 3.14

	const (
		ime_           = "Janko"
		plata          = 5000
		datum_rodjenja = "20.12.2002"
	)

	var prezime string

	chack := true

	fmt.Println(chack)
	
	fmt.Println(Pi)
	// fmt.Println(ime_, plata, datum_rodjenja)
}
