package euler

import (
	"sort"
)

func PrimesAndComposites(limit int) (primes map[int]bool, composites map[int]bool) {
	composites = make(map[int]bool)
	primes = make(map[int]bool, 0)
	p := 2

	for p < limit {
		// p is a prime
		primes[p] = true

		// all multiples of p are not primes
		for i := 2 ; i * p < limit ; i++ {
			composites[i*p] = true
		}

		// find the next number not yet marked as a composite
		p++
		for present := composites[p] ; present && p < limit ; present = composites[p] {
			p++
		}
	}

	return primes, composites
}

func PrimesUnderN(limit int) []int {
	primes, _ := PrimesAndComposites(limit)
	keys := make([]int, 0, len(primes))
	for key := range primes {
		keys = append(keys, key)
	}
	sort.Ints(keys)
	return keys
}

func NPrimes(n int) []int {
	primes := make([]int, 1, n)
	primes[0] = 2 // 2 is a prime

	attempt := 3 // start from 3

	for len(primes) < n {
		if AllInt(func (p int) bool { return attempt % p != 0}, primes) {
			primes = append(primes, attempt)
		}
		attempt += 2
	}

	return primes
}
