package main


import "fmt"


// Instructions
// Note: This exercise is a continuation of the need-for-speed exercise.

// In this exercise you'll be organizing races between various types of remote controlled cars. Each car has its own speed and battery drain characteristics.

// Cars start with full (100%) batteries. Each time you drive the car using the remote control, it covers the car's speed in meters and decreases the remaining battery percentage by its battery drain.

// If a car's battery is below its battery drain percentage, you can't drive the car anymore.

// The remote controlled car has a fancy LED display that shows two bits of information:

// The total distance it has driven, displayed as: "Driven <METERS> meters".
// The remaining battery charge, displayed as: "Battery at <PERCENTAGE>%".
// Each race track has its own distance. Cars are tested by checking if they can finish the track without running out of battery.

// 1. Drive the car
// Implement the Drive method on the Car that updates the number of meters driven based on the car's speed, and reduces the battery according to the battery drainage:

// speed := 5
// batteryDrain := 2
// car := NewCar(speed, batteryDrain)
// car.Drive()
// // car is now Car{speed: 5, batteryDrain: 2, battery: 98, distance: 5}
// Note: You should not try to drive the car if doing so will cause the car's battery to be below 0.

// 2. Display the distance driven
// Implement a DisplayDistance method on Car to return the distance as displayed on the LED display as a string:

// speed := 5
// batteryDrain := 2
// car := NewCar(speed, batteryDrain)

// fmt.Println(car.DisplayDistance())
// Output: "Driven 0 meters"
// 3. Display the battery percentage
// Implement the DisplayBattery method on Car to return the battery percentage as displayed on the LED display as a string:

// speed := 5
// batteryDrain := 2
// car := NewCar(speed, batteryDrain)

// fmt.Println(car.DisplayBattery())
// Output: "Battery at 100%"
// 4. Check if a remote control car can finish a race
// To finish a race, a car has to be able to drive the race's distance. This means not draining its battery before having crossed the finish line. Implement the CanFinish method that takes a trackDistance int as its parameter and returns true if the car can finish the race; otherwise, return false:

// speed := 5
// batteryDrain := 2
// car := NewCar(speed, batteryDrain)

// trackDistance := 100

// car.CanFinish(trackDistance)
// => true

type Car struct {
	battery, batteryDrain, speed, distance int
}

func NewCar(speed, batteryDrain int) Car {
	return Car{
		battery:      100,
		speed:        speed,
		batteryDrain: batteryDrain,
	}
}


func (car *Car)Drive(){
	if car.battery >= car.batteryDrain {
		car.battery -= car.batteryDrain
		car.distance += car.speed
	}
}

func (car *Car)DisplayDistance()string{
	return fmt.Sprintf("Driven %d meters", car.distance)
}

func (car *Car)DisplayBattery()string{
	return fmt.Sprintf("Battery at %d%%", car.battery)
}


func (car *Car) CanFinish(trackDistance int) bool{
	return trackDistance <= car.battery/car.batteryDrain*car.speed
}

func main() {
	Car := NewCar(5, 2)
	Car.Drive()
	Car.Drive()
	Car.Drive()

	fmt.Println(Car.DisplayDistance())
	fmt.Println(Car.DisplayBattery())
	fmt.Println(Car.CanFinish(100))

}
