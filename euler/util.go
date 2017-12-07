package euler

func SumInts(source []int) int {
	total := 0
	for _, i := range(source) {
		total += i
	}
	return total
}

