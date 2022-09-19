package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

type Text string
type Password string
type Price float64
type ID uint

type User struct {
	FirstName Text
	Password  Password
}

type Developer struct {
	id       ID
	language Text
}

func main() {
	p := Person{"Ime", "Prezime", 19}

	fmt.Printf("value: %v, type: %T \n", p, p)

	p2 := &Person{"Ime2", "Prezime2", 20}
	fmt.Printf("value: %v, type: %T \n", p2, p2)

	fmt.Println(p.FirstName)
	fmt.Println(p2.FirstName)

	fmt.Printf("%T", p2.FirstName) //Sting

	//====================================================================================

	p3 := Person{}

	fmt.Println(p3)

	p3.FirstName = "Name3"

	fmt.Println(p3)

	// pokazivac := &13
	// fmt.Println(pokazivac)

	//====================================================================================

	cijena := Price(21.51)

	fmt.Printf("%T %v \n", cijena, cijena)

	//=====================================================================================

	local_struct := struct {
		LastName  string
		FirstName string
		cijena    Price
	}{"Ime", "Pero", 12125.2}

	fmt.Printf("value: %v, type: %T \n\n", local_struct, local_struct)

	local_struct.cijena = 1.01

	fmt.Printf("value: %v, type: %T \n\n", local_struct, local_struct)

	//=======================================================================================

	person1 := struct {
		User
		Developer
	}{User: User{"Janko7262", "654:D14@sadasd[[g]fs"}, Developer: Developer{1, "Golang"}}

	fmt.Printf("Type: %T, Value: %v \n\n", person1, person1)

	fmt.Println(person1.User.Password) //radiii! :D

	//========================================================================================
}
