package main

import "fmt"

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	return map[string]int{
		"quarter_of_a_dozen": 3,
		"half_of_a_dozen":    6,
		"dozen":              12,
		"small_gross":        120,
		"gross":              144,
		"great_gross":        1728,
	}
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	return map[string]int{}
}

func AddItem(bill, units map[string]int, item, unit string) bool {

	value, my_error := units[unit]
	if !my_error {
		return false
	}

	bill[item] = value

	return true
}

func main() {
	bill := NewBill()
	units := Units()
	ok := AddItem(bill, units, "carrot", "dozen")
	fmt.Println(ok)
}
