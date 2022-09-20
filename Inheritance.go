package main

import "fmt"

type Entity struct {
	Id  uint64
	Age uint8
}

type Human struct {
	*Entity
	name string
}

type Animal struct {
	Entity
	Species string
}

func main() {

	human := Human{&Entity{21516565165, 21}, "Pero"}
	animal := Animal{Entity{654654654, 4}, "Lav"}

	animal2 := Animal{}

	animal2.Species, animal2.Id, animal2.Age = "Stevo", 32164, 41

	fmt.Printf("%v\n %v\n", human.Age, animal.Age)

	fmt.Printf("%v", animal2)
}
