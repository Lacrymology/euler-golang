package euler

import (
	"math"
)

func LargestPrime(n int) int {
	max := int(math.Sqrt(float64(n))) + 1
	primes := Primes(max)

	for _, p := range primes {
		if n % p != 0 {
			continue
		}

		for n % p == 0 && n > 1 {
			n = n / p
		}
		if n <= p {
			return p
		}
	}

	return n
}
