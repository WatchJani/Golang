package main

import "fmt"

type Speed int
type Text string

type Vehicle struct {
	MaxSpeed Speed
	Name     Text
}

func main() {

	my_car := Vehicle{152, "Golf 2"}
	my_car.MaxSpeed = 5
	fmt.Println(my_car.MaxSpeed, " ", my_car.Name)

	my_new_car := struct {
		MaxSpeed Speed
		Name     Text
	}{210, "BMW"}

	fmt.Println(my_new_car.MaxSpeed)
}
