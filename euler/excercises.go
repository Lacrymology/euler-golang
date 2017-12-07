package euler

func Ex001 () int {
	return SumInts(MultiplesOf3Or5(1000))
}

func Ex002 () int {
	return SumInts(FilterInts(Fib(4e6), func (i int) bool {
		return i % 2 == 0
	}))
}

func Ex003 () int {
	return LargestPrime(600851475143)
}
