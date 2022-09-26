package main

import "fmt"

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, time int) int {
	if time == 0 {
		return len(layers) * 2
	}
	return len(layers) * time
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (int, float64) {
	var noodles int
	var sauce float64

	for _, layer := range layers {
		switch layer {
		case "noodles":
			noodles++
		case "sauce":
			sauce++
		}
	}
	return noodles, sauce
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendsList, myList []string) {
	myList[len(myList)-1] = friendsList[len(friendsList)-1]
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, scaledQuantities int) []float64 {
	FloatScaledQuantities := float64(scaledQuantities) / 2

	var new_quantities []float64

	for _, quantitie := range quantities {
		new_quantities = append(new_quantities, quantitie*FloatScaledQuantities)
	}

	return new_quantities
}

func main() {
	fmt.Println(PreparationTime([]string{"sauce", "noodles", "sauce", "meat", "mozzarella", "noodles"}, 3))
	fmt.Println(Quantities([]string{"sauce", "noodles", "sauce", "meat", "mozzarella", "noodles"}))

	friendsList := []string{"noodles", "sauce", "mozzarella", "kampot pepper"}
	myList := []string{"noodles", "meat", "sauce", "mozzarella", "?"}

	AddSecretIngredient(friendsList, myList)

	fmt.Println(myList)

	quantities := []float64{1.2, 3.6, 10.5}

	fmt.Println(ScaleRecipe(quantities, 5))
}
