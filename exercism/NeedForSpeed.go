package main

import "fmt"

type Car struct {
	battery      int
	batteryDrain int
	speed        int
	distance     int
}

// NewCar creates a new remote controlled car with full battery and given specifications.
func NewCar(speed, batteryDrain int) Car {
	return Car{
		battery:      100,
		speed:        speed,
		batteryDrain: batteryDrain,
		// ne moram deklarisati je je pocetna vrijednost i onako 0(distance:     0,)
	}
} //Mogu predefinisati vrijednosti strukuture

// TODO: define the 'Track' type struct
type Track struct {
	distance int
}

// NewTrack creates a new track
func NewTrack(distance int) Track {
	return Track{
		distance: distance,
	}
}

// Drive drives the car one time. If there is not enough battery to drive one more time,
// the car will not move.
func Drive(car Car) Car {
	if car.battery >= car.batteryDrain {
		car.battery -= car.batteryDrain
		car.distance += car.speed
		return car
	}
	return car

	// if car.battery < car.batteryDrain {
	// 	return car
	// }

	// return Car{
	// 	batteryDrain: car.batteryDrain,
	// 	distance:     car.distance + car.speed,
	// 	battery:      car.battery - car.batteryDrain,
	// 	speed:        car.speed,
	// }
}

// CanFinish checks if a car is able to finish a certain track.
func CanFinish(car Car, track Track) bool {
	return track.distance <= car.battery/car.batteryDrain*car.speed
}

func main() {
	Car := NewCar(5, 5)
	Car = Drive(Car)
	Car = Drive(Car)
	Car = Drive(Car)
	Car = Drive(Car)
	Car = Drive(Car)
	Car = Drive(Car)
	Car = Drive(Car)
	Car = Drive(Car)
	Car = Drive(Car)
	Car = Drive(Car)
	Car = Drive(Car)
	Car = Drive(Car)
	Car = Drive(Car)
	Car = Drive(Car)
	Car = Drive(Car)
	Car = Drive(Car)
	Car = Drive(Car)
	Car = Drive(Car)
	Car = Drive(Car)
	Car = Drive(Car)
	Car = Drive(Car)

	distance := 100
	track := NewTrack(distance)

	fmt.Println(Car)
	fmt.Println(CanFinish(Car, track))
}
