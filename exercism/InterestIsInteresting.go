package main

import "fmt"

// InterestRate returns the interest rate for the provided balance.
func InterestRate(balance float64) float32 {
	switch {
	case 0 <= balance && balance < 1000:
		return 0.5
	case 1000 <= balance && balance < 5000:
		return 1.621
	case 5000 <= balance:
		return 2.475
	default:
		return 3.213
	}
}

// Interest calculates the interest for the provided balance.
func Interest(balance float64) float64 {
	return balance / 100 * float64(InterestRate(balance))
}

// AnnualBalanceUpdate calculates the annual balance update, taking into account the interest rate.
func AnnualBalanceUpdate(balance float64) float64 {
	return balance + Interest(balance)
}

// YearsBeforeDesiredBalance calculates the minimum number of years required to reach the desired balance.
func YearsBeforeDesiredBalance(balance, targetBalance float64) int {
	var counter int

	for balance <= targetBalance {
		counter++
		balance = AnnualBalanceUpdate(balance)
	}

	return counter
}

func main() {
	fmt.Println(InterestRate(200))
	fmt.Println(Interest(200.75))
	fmt.Println(AnnualBalanceUpdate(200.75))
	fmt.Println(YearsBeforeDesiredBalance(200.75, 214.88))
}
