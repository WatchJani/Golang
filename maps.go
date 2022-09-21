package main

import (
	"fmt"
	"reflect"
)

func main() {

	var my_map = map[string]int{}

	fmt.Printf("lenght: %v\n", len(my_map))
	fmt.Printf("value: %v\n", my_map)
	fmt.Printf("type: %T\n", my_map)

	//=================================================================

	my_map["name_of_key"] = 1
	fmt.Println(my_map)

	//=================================================================

	my_map = map[string]int{
		"Golang": 1,
		"C":      3,
		"C++":    2,
	}
	fmt.Println(my_map)

	//==================================================================

	m1 := make(map[string]int32)
	m2 := map[string]int32{}

	fmt.Println(reflect.DeepEqual(m1, m2))

	//==================================================================

	m3 := make(map[string]int32, 2)
	m4 := map[string]int32{
		"Janko": 10,
	}

	fmt.Printf("%v: Len=%v Values=%v\n", "m3", len(m3), m3)
	fmt.Printf("%v: Len=%v Values=%v\n", "m4", len(m4), m4)

	m3 = map[string]int32{
		"Mark":  20,
		"Sundy": 40,
		"Nick":  101,
	}

	fmt.Printf("%v: Len=%v Values=%v\n", "m3", len(m3), m3)

	//===================================================================

	n1 := map[string]int32{
		"Mark":  1,
		"Sundy": 30,
		"Nick":  1041,
	}

	value, _ := n1["Mark"]

	fmt.Printf("Value: %v\n", value)

	_, chack := n1["Mark"]

	fmt.Printf("Is exist: %v\n", chack)

	value2, chack2 := n1["Sasd"]

	fmt.Printf("Value: %v, Is exist: %v\n\n", value2, chack2)

	//===================================================================

	n1["Mark"] *= 10
	fmt.Printf("Value: %v\n", n1)

	//===================================================================

	delete(n1, "Mark")
	fmt.Printf("Value: %v\n", n1)

	//===================================================================

	n1["Janko"] = 10
	fmt.Printf("Value: %v\n", n1)

	//===================================================================
}
