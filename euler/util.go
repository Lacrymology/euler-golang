package euler

func SumInts(source <-chan int) int {
	total := 0
	for i := range(source) {
		total += i
	}
	return total
}

//func FilterInts(source []int, f func (int) bool) []string {
//}
