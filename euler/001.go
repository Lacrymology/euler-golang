package euler

func MultiplesOf3Or5(limit int) <-chan int {
	ret := make(chan int)

	go func () {
		for i := 1 ; i < limit ; i++ {
			if i % 3 == 0 || i % 5 == 0 {
				ret <- i
			}
		}
		close(ret)
	}()

	return ret
}
