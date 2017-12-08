package euler

func SmallestMultiple(limit int) int {
	values := make([]int, limit)
	for i := 0 ; i < limit ; i++ {
		values[i] = i+1
	}
	return ReduceInt(LCM, values, 1)
}
