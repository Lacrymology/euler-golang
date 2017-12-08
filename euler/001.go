package euler

func MultiplesOf3Or5(limit int) []int {
	ret := make([]int, 0)

	for i := 1; i < limit; i++ {
		if i%3 == 0 || i%5 == 0 {
			ret = append(ret, i)
		}
	}

	return ret
}
