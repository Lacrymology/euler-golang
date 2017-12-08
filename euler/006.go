package euler

func SquareDiff(limit int) int {
	var squareSum, sumSquared int

	for i := 1 ; i <= limit ; i++ {
		squareSum += i*i
		sumSquared += i
	}
	sumSquared *= sumSquared

	return sumSquared - squareSum
}
