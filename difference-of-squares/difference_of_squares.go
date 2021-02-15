package diffsquares

// SquareOfSum gets int and return int
func SquareOfSum(n int) int {
	sum := (n * (n + 1)) / 2
	return sum * sum
}

// SumOfSquares gets int and return int
func SumOfSquares(n int) int {
	return (n * (n + 1) * ((n * 2) + 1)) / 6
}

// Difference gets int and return int
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
