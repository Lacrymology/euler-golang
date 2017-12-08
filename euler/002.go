package euler

func Fib(limit int) []int {
	ret := make([]int, 0)

	i, j := 1, 1

	for j < limit {
		ret = append(ret, j)
		i, j = j, i+j
	}

	return ret
}
