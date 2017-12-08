package euler

func ReduceInt(r func(int, int) int, vals []int, initial int) int {
	for _, v := range vals {
		initial = r(initial, v)
	}
	return initial
}

func AllInt(p func(int) bool, vals []int) bool {
	for _, i := range vals {
		if !p(i) {
			return false
		}
	}
	return true
}

func AnyInt(p func(int) bool, vals []int) bool {
	return AllInt(func(i int) bool { return !p(i) }, vals)
}

func SumInts(source []int) int {
	return ReduceInt(func(a, b int) int { return a + b }, source, 0)
}
