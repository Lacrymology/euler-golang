package euler

import (
	"fmt"
	"math"
)

func LargestPalindrome(figures int) int {
	a := int(math.Pow10(figures)) - 1
	b := a

	max := 0

	for a > 99 {
		for b > 99 {
			x := a * b
			//fmt.Printf("%d * %d: %d\n", a, b, x)
			if IsPalindrome(fmt.Sprintf("%d", x)) && x > max {
				max = x
			}
			b--
		}
		a--
		b = a
	}
	return max
}
