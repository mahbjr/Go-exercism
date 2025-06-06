package diffsquares

// SquareOfSum calculates the square of the sum of first n natural numbers
// Formula: (n*(n+1)/2)^2
func SquareOfSum(n int) int {
	sum := n * (n + 1) / 2
	return sum * sum
}

// SumOfSquares calculates the sum of squares of first n natural numbers
// Formula: n*(n+1)*(2n+1)/6
func SumOfSquares(n int) int {
	return n * (n + 1) * (2*n + 1) / 6
}

// Difference calculates the difference between square of sum and sum of squares
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
