package euler

func FilterInts(source []int, f func(int) bool) []int {
	dest := make([]int, 0)
	for _, v := range source {
		if f(v) {
			dest = append(dest, v)
		}
	}
	return dest
}

func IsPalindrome(input string) bool {
	for i := 0; i < len(input)/2; i++ {
		if input[i] != input[len(input)-i-1] {
			return false
		}
	}
	return true
}

func GCD(a, b int) int {
	for a != b {
		switch {
		case a > b:
			a -= b
		case b > a:
			b -= a
		}
	}
	return a
}

func LCM(a, b int) int {
	return a * b / GCD(a, b)
}
