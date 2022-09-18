package main

import "fmt"

func main() {
	number := float64(200.15516)

	fmt.Println((number - float64(int(number))))

	fmt.Println(number - 200.0) //0.15515999999999508 zasto nije 0.15516

	fmt.Printf("%.5f", number) //sad je dobro
}
