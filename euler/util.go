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
