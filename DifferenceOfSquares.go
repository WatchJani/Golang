// Instructions
// Find the difference between the square of the sum and the sum of the squares of the first N natural numbers.

// The square of the sum of the first ten natural numbers is (1 + 2 + ... + 10)² = 55² = 3025.

// The sum of the squares of the first ten natural numbers is 1² + 2² + ... + 10² = 385.

// Hence the difference between the square of the sum of the first ten natural numbers and the sum of the squares of the first ten natural numbers is 3025 - 385 = 2640.

// You are not expected to discover an efficient solution to this yourself from first principles; research is allowed, indeed, encouraged. Finding the best algorithm for the problem is a key skill in software engineering.


package main

import "fmt"

func SquareOfSum(n int) int {
	sum := (n * (n + 1)) / 2
	return sum * sum
}

// https://www.vedantu.com/question-answer/the-sum-of-squares-of-first-ten-natural-numbers-class-11-maths-cbse-5fbc82800736f77891d87306
func SumOfSquares(n int) int {
	return (n * (n + 1) * (2*n + 1)) / 6
}

func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}

func main() {
	fmt.Print(Difference(10))
}
