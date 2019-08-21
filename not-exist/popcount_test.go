package not_exist_test

// the idea is simple, each time we set last 1 bit to 0 until we got a 0
func kernighan(x int) int {
	var count int
	for ; x != 0; x &= (x - 1) {
		count++
	}
	return count
}
