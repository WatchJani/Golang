package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()

	var niz [2]string
	fmt.Println(niz)

	niz[0] = "Janko"
	niz[1] = "Kondic"

	fmt.Println(len(niz))

	var NewArray [5]int = [5]int{
		2,
		54,
		54,
		8,
		47654654654,
	}

	MyNewArray := [5]float64{5}

	fmt.Println(MyNewArray)

	fmt.Println(NewArray[len(NewArray)-1])

	my_Mega_List := [...]string{
		"C++",
		"Golang",
		"C#",
	}

	fmt.Println(len(my_Mega_List))

	fmt.Println(t)

	//==========================================================================================
	//PONAVLJANJE NAUCENOG
	//Sintaksas

	NewArray1 := [...]int{
		1,
		2,
		2,
	}

	Languages := [...]string{
		"Golan",
		"C++",
		"C",
		"JavaScript",
		"C#",
	}

	fmt.Println(NewArray1[1 : len(NewArray1)-1])
	fmt.Println(Languages[len(Languages)-2:])

	x := [5]int{0: 5, 4: 5}

	fmt.Println(x) //RADIII!!!!
}
