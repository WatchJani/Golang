package main

import "fmt"

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	var sum int

	for _, birds := range birdsPerDay {
		sum += birds
	}

	return sum
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	// var sum int
	offset := 7 * (week - 1)
	// for i := week*DaysInTheWeek - DaysInTheWeek; i <= week*DaysInTheWeek-1; i++ {
	// 	sum += birdsPerDay[i]
	// }

	return TotalBirdCount(birdsPerDay[offset : offset+7])
}

func FixBirdCountLog(birdsPerDay []int) []int {
	for i := 0; i < len(birdsPerDay); i += 2 {
		birdsPerDay[i] += 1
	}

	return birdsPerDay
}

func main() {
	fmt.Println(TotalBirdCount([]int{2, 5, 0, 7, 4, 1, 3, 0, 2, 5, 0, 1, 3, 1}))
	fmt.Println(BirdsInWeek([]int{2, 5, 0, 7, 4, 1, 3, 0, 2, 5, 0, 1, 3, 1}, 2))
	fmt.Println(FixBirdCountLog([]int{2, 5, 0, 7, 4, 1}))
}
