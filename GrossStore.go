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

	bill[item] += value

	return true
}


// RemoveItem removes an item from customer bill.

//bill je mapa sa stvarima koje je kupio 
//units je mapa koja ima sve sto se moze kupiti
//item je sta bill kupuje
//unit 

func RemoveItem(bill, units map[string]int, item, unit string) bool {
	bill_pocket, error_itme := bill[item]
	value, error_unit := units[unit]


	if !error_itme || !error_unit || value > bill_pocket{
		return false
	}else if value == bill_pocket{
		delete(bill, item)
	}else{
		bill[item] -= value
	}
	
	return true
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	value, set_error := bill[item]
	return value, set_error
}

func main() {
	bill := NewBill()
	units := Units()
	ok := RemoveItem(bill, units, "carrot", "dozen")
	fmt.Println(ok)
}
