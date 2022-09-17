package main

import (
	"fmt"
	"reflect"
)

func main() {

	number := 132
	my_string := "Janko"

	my_float := 500.45451
	my_special_float := float64(number)

	fmt.Printf("%v\n", number) //ispisuje vrijednost varijable
	fmt.Printf("%T\n", number) //ispisuje vrijednost
	fmt.Printf("%v\n\n", my_string)
	fmt.Printf("%q\n", my_string) //dodaje navodnike

	fmt.Printf("%f %f \n", my_float, my_special_float)

	if user := true; user {
		//testiram lokaciju varijable
		fmt.Printf("Vrijednost varijable(%T) je (%.2f), koja se nalazi na lokaciji %v \n\n", my_special_float, my_special_float, &my_special_float)
	}

	fmt.Println(`moze li
	novi red???`)

	fmt.Printf("|%6s|%-10.2f|\n", my_string, my_special_float)

	fmt.Printf("%v\n", reflect.TypeOf(number))

	fmt.Printf("%s \n", " \bcao svima")
	fmt.Printf("\\%s \n", "moze sad")           //dodaje \(mada ovjde moze ici bilo koji znak)
	fmt.Printf("\t%s \n", "cao")                //tab dodaje
	fmt.Printf("\v%-0.2f \n", my_special_float) // ovo ne radi?????
	// fmt.Printf("\' %s\n", my_string)
	fmt.Printf(" \bMoze") //??????
}
