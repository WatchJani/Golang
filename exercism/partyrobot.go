package main

import "fmt"

func HappyBirthday(name string, age int) string {
	// your_age := fmt.Sprint(age)
	return fmt.Sprintf("Happy birthday %v! You are now %v years old!", name, age)
}

func Welcome(name string) string {
	return fmt.Sprintf("Welcome to my party, %v!", name)
}

func AssignTable(name string, table int, neighbor, direction string, distance float64) string {
	welcomeMsg := Welcome(name)
	tableAssignMsg := fmt.Sprintf("You have been assigned to table %03d.", table)
	tableDirMsg := fmt.Sprintf("Your table is %s, exactly %.1f meters from here.", direction, distance)
	neighborMsg := fmt.Sprintf("You will be sitting next to %s.", neighbor)
	return fmt.Sprintf("%s\n%s %s\n%s", welcomeMsg, tableAssignMsg, tableDirMsg, neighborMsg)
}

// Welcome to my party, Christiane!
// You have been assigned to table 027. Your table is on the left, exactly 23.8 meters from here.
// You will be sitting next to Frank.

func main() {
	fmt.Println(Welcome("Janko"))
	fmt.Println(HappyBirthday("Janko", 20))
	fmt.Println(AssignTable("Christiane", 22, "Frank", "straight ahead", 9.2394381))
}
