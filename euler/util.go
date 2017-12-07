package euler

func SumInts(source []int) int {
	total := 0
	for _, i := range(source) {
		total += i
	}
	return total
}

func FilterInts(source []int, f func (int) bool) []int {
	dest := make([]int, 0)
	for _, v := range(source) {
		if f(v) {
			dest = append(dest, v)
		}
	}
	return dest
}

func Primes(limit int) []int {
	composites := make(map[int]bool)
	primes := make([]int, 0)
	p := 2

	for p < limit {
		// p is a prime
		primes = append(primes, p)

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

	return primes
}
