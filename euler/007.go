package euler

func NthPrime(n int) int {
	primes := NPrimes(n)
	return primes[len(primes)-1]
}
