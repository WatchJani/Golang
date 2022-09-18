package main

import "fmt"

func Println(x int) int {
	return x * 50
}

func IntString(name, age string) (string, string) {
	return name, age
}

func main() {
	fmt.Println(Println(5))
	name, age := IntString("Janko", "19")

	fmt.Println(name, age)
}
