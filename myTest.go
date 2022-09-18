package main

import (
	"fmt"
	"strings"
)

func WelcomeMessage(customer string) string {
	return "Welcome to the Tech Palace," + customer
}

func CleanupMessage(oldMsg string) string {
	return strings.TrimSpace(strings.ReplaceAll(oldMsg, "*", "")) 
}

func main() {
	name := "Janko"

	message := `**************************
	*    BUY NOW, SAVE 10%   *
	**************************`

	fmt.Println(WelcomeMessage(name))
	fmt.Println(CleanupMessage(message))

}
