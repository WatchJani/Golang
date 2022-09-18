package main

import (
	"fmt"
	"time"
)

func main() {
	ime := "Janko"
	prezime := "Janko"

	//mogu se porediti stringovi :D
	if ime == prezime {
		fmt.Printf(ime + prezime)
	}

	fmt.Printf("\n%T \n", ime[0]) //uint8

	// date and time

	t := time.Now()

	fmt.Println(t)

	fmt.Printf("%T \n", t)
	fmt.Println(t.Location())

	current_time := time.Now()
	previous_time := current_time.Format("02-01-2006")

	fmt.Printf("%T\n", previous_time) //string

	if broj := "5"; current_time.Format("02") == broj {
		fmt.Println("Vaš broj odgovara danu u mjesecu")
	} else {
		if broj < current_time.Format("02") {
			// fmt.Println(current_time.Format("02") - broj)
		}
	}

	fmt.Println(previous_time)

	new_time := time.Now()

	fmt.Println(new_time)

	timeInNewYork, _ := time.LoadLocation("America/New_York")

	fmt.Printf("%s %s", timeInNewYork, new_time.In(timeInNewYork))

	fmt.Printf("\n\n%s\n\n", new_time.In(timeInNewYork).Location())

	location, _ := time.LoadLocation("America/New_York")
	now := time.Now().In(location)

	fmt.Println(now)

	my_date := time.Date(2022, 12, 20, 18, 59, 1, 321, time.UTC)

	next_date := my_date.AddDate(1, 0, 0)

	fmt.Println(my_date)
	fmt.Println(next_date)

	fmt.Println(my_date.Add(time.Minute))


	fmt.Println(my_date.Format(time.RubyDate))
}
