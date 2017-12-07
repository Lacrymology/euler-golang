package euler

import (
	"fmt"
)

func Ex001 () {
	muls := MultiplesOf3Or5(1000)
	fmt.Println(SumInts(muls))
}
